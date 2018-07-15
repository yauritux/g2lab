package main

import (
	"fmt"
	"math"
)

// I interface
type I interface {
	M()
}

// T struct
type T struct {
	S string
}

// T pointer receiver method
// Implicitly implement I interface
func (t *T) M() {
	fmt.Println(t.S)
}

// F float64
type F float64

// F pointer receiver method
// Implicitly implement I interface
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%+v, %T)\n", i, i)
}
