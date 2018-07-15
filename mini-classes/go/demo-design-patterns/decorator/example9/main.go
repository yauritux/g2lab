package main

import (
	"fmt"
	"log"
	"os"
)

type Echoer interface {
	Echo() error
	Dump() string
}

type text struct {
	s string
}

func (t text) Echo() error {
	_, err := fmt.Println(t.s)
	return err
}

func (t text) Dump() string {
	return fmt.Sprintf(t.s)
}

func printIt(e Echoer) {
	e.Echo()
}

func dumpIt(e Echoer) string {
	return e.Dump()
}

// Decorates the Echoer interface
func loggingDecorator(e Echoer, l *log.Logger) Echoer {
	l.Println(">>>", dumpIt(e))
	return e
}

func main() {
	t := text{s: "Hello Gopher!"}

	// not decorated, sends t, which has method Echo, which implements Echoer interface method
	printIt(t)

	myLogger := log.New(os.Stdout, "###", 3)
	td := loggingDecorator(t, myLogger)
	td.Echo()
}
