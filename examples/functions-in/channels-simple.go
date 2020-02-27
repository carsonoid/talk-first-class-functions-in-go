package main

import "fmt"

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

func main() {
	// Make a channel that can buffer one print func at a time
	printerChan := make(chan PrintFunc, 1)

	// Create a function and assign it to the variable "printerOne"
	printerOne := func(msg string) { fmt.Println("printerOne says:", msg) }

	// Send the function variable to the channel
	printerChan <- printerOne

	// Get the printer function off the channel and use it
	p := <-printerChan
	p("Hello")

	// Send a function to the channel without assigning it.
	printerChan <- func(msg string) { fmt.Println("printerTwo says:", msg) }

	// Get and print without assigment
	(<-printerChan)("Howdy")
}
