package main

import "fmt"

func main() {
	a := [2]string{"Hello", "Gopher"}
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
