package main

import "fmt"

type Greeter interface {
	SayHello()
}

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

// PrintFunc implements the Greeter interface
func (pf PrintFunc) SayHello() {
	fmt.Println("Hello from a method of a function")
}

func main() {
	p := new(PrintFunc)
	p.SayHello()
}
