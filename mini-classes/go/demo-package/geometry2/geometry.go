package main

import (
	"fmt"
	"log"

	"github.com/yauritux/g2lab/mini-classes/go/demo-package/rectangle"
)

var rectLength, rectWidth float64 = 6, 7

func init() {
	fmt.Println("main package initialized")
	if rectLength < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	fmt.Printf("Area of the rectangle %.2f\n", rectangle.Area(rectLength, rectWidth))
	fmt.Printf("Diagonal of the rectangle %.2f\n", rectangle.Diagonal(rectLength, rectWidth))
}
