package main

import (
	"fmt"
	"strings"
)

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

func newStatefulPrinter() PrintFunc {
	data := []string{}
	return func(msg string) {
		data = append(data, msg)
		fmt.Println(strings.Join(data, " "))
	}
}

func main() {
	sayAndRemember := newStatefulPrinter()
	sayAndRemember("Hello")
	sayAndRemember("I have")
	sayAndRemember("state!")

	sayAndRemember = newStatefulPrinter()
	sayAndRemember("I'm new")
	sayAndRemember("and have")
	sayAndRemember("new state!")
}
