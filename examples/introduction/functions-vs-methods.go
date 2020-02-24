package main

import (
	"fmt"
)

func main() {
	printMsg("Printed with the 'printMsg' function")

	p := Printer{}
	p.printMsg("Printed with the 'printMsg' method owned by 'p'")

	pfunc := func(msg string) { fmt.Println(msg) }
	pfunc("Printed with an anonymous function assigned to a variable")

	func(msg string) {
		fmt.Println(msg)
	}("Printed with an anonymous function without assignment")
}

// printMsg is a simple function. It belongs to the current package
func printMsg(msg string) {
	fmt.Println(msg)
}

// Printer is an empty struct which owns the printMsg method
type Printer struct{}

// printMsg is technically a method since it belongs to a particular "Printer"
func (p *Printer) printMsg(msg string) {
	fmt.Println(msg)
}
