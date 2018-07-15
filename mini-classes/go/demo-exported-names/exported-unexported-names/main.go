package main

import (
	"fmt"

	"github.com/yauritux/demo-exported-unexported/test/animals"
	"github.com/yauritux/demo-exported-unexported/test/counters"
)

func main() {
	// Create a variable of the exported type and
	// initialize the value to 10.
	counter := counters.NewAlertCounter(10)

	fmt.Printf("Counter: %d\n", counter)

	// Create an object of type Dog from the animals package.
	/*
		dog := animals.Dog{
			Animal: animals.Animal{
				Name: "Chole",
				Age:  1,
			},
			BarkStrength: 10,
		}
	*/
	dog := animals.Dog{
		BarkStrength: 10,
	}
	dog.Name = "Chole"
	dog.Age = 1

	fmt.Printf("Dog: %#v\n", dog)
}
