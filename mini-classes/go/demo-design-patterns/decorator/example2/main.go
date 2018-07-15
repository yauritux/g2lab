package main

import (
	"fmt"
	"math/rand"
	"time"
)

func timeIt(fn func(s string)) func(string) {
	return func(s string) {
		start := time.Now()
		rand.Seed(rand.Int63n(1000000))
		r := rand.Int63n(10)
		fmt.Println("sleeping:", r)
		time.Sleep(time.Duration(r) * time.Second)
		defer func() {
			fmt.Printf("[timeIt] took: %v\n", time.Since(start))
		}()
		fn(s)
	}
}

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	//without decorator
	printSomething("Yauri")

	//with decorator
	timeIt(printSomething)("Yauri")
}
