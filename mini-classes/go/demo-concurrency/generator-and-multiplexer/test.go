package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("I'm all ears...")
	c := fanIn(boring("Joe:"), boring("Ann:"))
	for i := 0; i <= 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// fanIn is a multiplexer that combine 2 channels into one particular channel
func fanIn(channel1, channel2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-channel1
		}
	}()
	go func() {
		for {
			c <- <-channel2
		}
	}()
	return c
}
