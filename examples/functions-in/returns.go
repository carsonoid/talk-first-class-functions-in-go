package main

import "fmt"

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

func getPrinter() PrintFunc {
	return func(msg string) {
		fmt.Println(msg)
	}
}

func main() {
	p := getPrinter()
	p("Hello from the function returned by getPrinter")
}
