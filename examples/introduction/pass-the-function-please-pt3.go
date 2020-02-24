package main

import (
	"fmt"
)

func main() {
	printWith(fmt.Println, "Printed with fmt's Println function")

	printWith(printMsg, "Printed with the 'printMsg' function")

	p := new(Printer)
	printWith(p.printMsg, "Printed with the Printer's 'printMsg' method")

	pfunc := func(msgs ...interface{}) (int, error) { return fmt.Println(msgs...) }
	printWith(pfunc, "Printed with an anonymous function assigned to a variable")

	printWith(
		func(msgs ...interface{}) (int, error) {
			return fmt.Println(msgs...)
		},
		"Printed with an anonymous function without assignment",
	)
}

// PrintFunc is any function or method that takes a string
type PrintFunc func(...interface{}) (int, error)

// printWith takes any function that takes a string and passes it a message
func printWith(p PrintFunc, msgs ...interface{}) {
	p(msgs...)
}

// printMsg is a simple wrapper around fmt.Println
func printMsg(msgs ...interface{}) (int, error) {
	return fmt.Println(msgs...)
}

// Printer is an empty struct which owns the printMsg method
type Printer struct{}

func (p *Printer) printMsg(msgs ...interface{}) (int, error) {
	return fmt.Println(msgs...)
}
