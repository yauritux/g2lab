package main

import (
	"fmt"
	"time"
)

func says(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go says("world")
	says("hello")
}
