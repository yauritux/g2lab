package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type somedata struct {
	s string
}

var counter int

func (d *somedata) doesAppThings() string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf(">>> %v", d.s)
}

// loggingDecorator outputs to stdout using the custom logger "l"
func loggingDecorator(f func() string, l *log.Logger) func() string {
	return func() string {
		l.Println("start...")
		defer func() { l.Println("end...") }()
		return f()
	}
}

// metricDecorator increments a global "counter" variable
func metricDecorator(f func() string, counter *int) func() string {
	*counter++
	return func() string {
		defer func() { fmt.Println("count", *counter) }()
		return f()
	}
}

func main() {
	s := somedata{s: "yoooloooo"}

	sNonDecorated := s.doesAppThings()
	fmt.Println("Non-Decorated:", sNonDecorated)

	myLogger := log.New(os.Stdout, "###", 3)

	// count = 1
	sDecorated := loggingDecorator(metricDecorator(s.doesAppThings, &counter), myLogger)()
	fmt.Println("decorated    ", sDecorated)

	// count = 2
	sDecorated = loggingDecorator(metricDecorator(s.doesAppThings, &counter), myLogger)()
	fmt.Println("decorated    ", sDecorated)

	// count = 3
	sDecorated = loggingDecorator(metricDecorator(s.doesAppThings, &counter), myLogger)()
	fmt.Println("decorated    ", sDecorated)
}
