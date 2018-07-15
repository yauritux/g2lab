package main

import "fmt"

type somedata struct {
	s string
}

func (d *somedata) printSomething(prefix string) {
	fmt.Println(prefix, d.s)
}

type DoPrintSomething func(string)

func decorate(f DoPrintSomething) DoPrintSomething {
	return func(s string) {
		fmt.Println("[decorate] before")
		f(s)
		fmt.Println("[decorate] after")
	}
}

func main() {
	d := somedata{s: "I'm not decorated\n"}
	d.printSomething(">>")

	e := somedata{s: "I'm now decorated... woohoo!"}
	decorate(e.printSomething)("==>")
}
