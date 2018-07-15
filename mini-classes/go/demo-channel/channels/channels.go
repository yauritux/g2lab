package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // sending sum to channel 'c'
}

func main() {
	start := time.Millisecond
	s := []int{5, 12, 7, -3, 1, 0, -2, 55, -25}
	c := make(chan int)

	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)
	x, y := <-c, <-c // receive from channel 'c'

	fmt.Printf("x=%d, y=%d, total=%d\n", x, y, x+y)
	end := time.Millisecond
	fmt.Printf("Time taken: %dms\n", (end - start))

}
