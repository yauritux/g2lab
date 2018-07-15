package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		//time.Sleep(time.Second)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func boringWithChannel(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func boringGenerator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	/*
		go boring("boring!")
		fmt.Println("I'm listening...")
		time.Sleep(2 * time.Second)
		fmt.Println("You're boring. I'm leaving.")
	*/
	/*
		c := make(chan string)
		go boringWithChannel("boring!", c)
		for i := 0; i < 5; i++ {
			fmt.Printf("You say: %q\n", <-c)
		}
		fmt.Println("You're boring; I'm leaving.")
	*/
	c := boringGenerator("boring!")
	for i := 0; i < 7; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
