package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Printf("x = %d\n", v.X)
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v)
}
