package main

import (
	"fmt"
)

func main() {
	printMsg("Printed with printMsg Directly!")

	printWith(printMsg, "Printed with the 'printMsg' function")

	p := new(Printer)
	printWith(p.printMsg, "Printed with the Printer's 'printMsg' method")

	pfunc := func(msg string) { fmt.Println(msg) }
	printWith(pfunc, "Printed with an anonymous function assigned to a variable")

	printWith(
		func(msg string) {
			fmt.Println(msg)
		},
		"Printed with an anonymous function without assignment",
	)
}

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

// printWith takes a PrintFunc and passes it a message
func printWith(p PrintFunc, msg string) {
	p(msg)
}

// printMsg is a simple wrapper around fmt.Println
func printMsg(msg string) {
	fmt.Println(msg)
}

// Printer is an empty struct which owns the printMsg method
type Printer struct{}

func (p *Printer) printMsg(msg string) {
	fmt.Println(msg)
}
