package main

import (
	"fmt"
)

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

var (
	packagePrinter = func(msg string) {
		fmt.Println(msg)
	}
)

func main() {
	packagePrinter("Printed with the packagePrinter anonymous function")
}
