package main

import "fmt"

type somedata struct {
	s string
}

func (d *somedata) printSomething() {
	fmt.Println(d.s)
}

func decorate(f func()) func() {
	return func() {
		fmt.Println("[decorator] before")
		f()
		fmt.Println("[decorator] after")
	}
}

func main() {
	//with decorator
	d := somedata{s: "I'm not decorated\n"}
	d.printSomething()

	//without decorator
	e := somedata{s: "I'm now decorated.... whoooo!"}
	decorate(e.printSomething)()
}
