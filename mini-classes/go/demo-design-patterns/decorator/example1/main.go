package main

import "fmt"

func decorate(fn func(s string)) func(string) {
	return func(s string) {
		fmt.Println("[decorator] before")
		fn(s)
		fmt.Println("[decorator] after")
	}
}

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	//without decorator
	printSomething("Yauri")

	//with decorator
	decorate(printSomething)("Yauri")
}
