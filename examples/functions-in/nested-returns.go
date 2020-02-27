package main

import (
	"fmt"
	"time"
)

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

func getPrinterMaker() func() PrintFunc {
	return func() PrintFunc {
		fmt.Printf("Making new printer: ")
		return func(msg string) {
			fmt.Println(msg)
		}
	}
}

func main() {
	makePrinter := getPrinterMaker()

	for i := 0; i < 3; i++ {
		p := makePrinter()
		time.Sleep(time.Second)
		p("Hello from a newly-generated printer")
		time.Sleep(time.Second)
	}
}
