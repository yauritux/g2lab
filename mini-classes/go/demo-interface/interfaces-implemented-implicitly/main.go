package main

import "fmt"

// I interface
type I interface {
	M()
}

// T struct type
type T struct {
	S string
}

// M is a T's method that implement I interface.
// This means type T implements the interface I.
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
