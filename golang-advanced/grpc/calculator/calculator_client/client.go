package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/yauritux/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)
	fmt.Printf("Created client: %v\n", c)

	doUnary(c)
	fmt.Println()
	doServerStreaming(c)
	fmt.Println()
	doErrorUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary server RPC...")
	req := &calculatorpb.Request{
		Num1: 3,
		Num2: 10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call sum RPC: %v\n", err)
	}

	fmt.Println(res)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeDecomposition Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecRequest{
		Number: 12,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeDecomposition Server RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res)
	}
}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a SquareRoot Unary RPC...")

	// correct call
	doErrorCall(c, 10)

	// error call
	doErrorCall(c, -2)
}

func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{
		Number: n,
	})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			//actual error from GRPC (user error)
			fmt.Println("Error message from server: ", respErr.Code(), "-", respErr.Message())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
			}
		} else {
			log.Fatalf("Big error calling SquareRoot: %v\n", err)
		}
	}
	fmt.Printf("Result of SquareRoot of %v: %v\n", n, res.GetNumberRoot())
}
