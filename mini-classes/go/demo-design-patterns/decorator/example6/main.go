package main

import (
	"fmt"
	"time"
)

type somedata struct {
	s string
}

func (d *somedata) returnStringWithPrefix(prefix string) string {
	return fmt.Sprintf("%s %s", prefix, d.s)
}

type DoPrintSomething func(string) string

func decorate(f DoPrintSomething) DoPrintSomething {
	return func(s string) string {
		start := time.Now()
		defer func() { fmt.Println("[decorate] took:", time.Since(start)) }()
		return f(s)
	}
}

func main() {
	d := somedata{"I'm not decorated\n"}
	ds := d.returnStringWithPrefix(">>")
	fmt.Println(ds)

	e := somedata{"I'm now decorated... woohoo!"}
	es := decorate(e.returnStringWithPrefix)("==>")
	fmt.Println(es)
}
