package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("I'm all ears...")
	c := boring("Joe")
	timeout := time.After(1 * time.Second)
	for {
		select {
		case <-c:
			fmt.Println(<-c)
		case <-timeout:
			fmt.Println("You're too slow")
			return
		}
	}
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
