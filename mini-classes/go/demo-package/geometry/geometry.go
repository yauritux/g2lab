package main

import (
	"fmt"

	"github.com/yauritux/g2lab/mini-classes/go/demo-package/rectangle"
)

func main() {
	fmt.Println("Geometrical shape properties")
	var rectLength, rectWidth float64 = 6, 7
	fmt.Printf("area of the rectangle %.2f\n", rectangle.Area(rectLength, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f\n", rectangle.Diagonal(rectLength, rectWidth))
}
