package main

import "fmt"

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

func callPrinters(msg string, p1, p2 PrintFunc) {
	p1("p1 says: " + msg)
	p2("p2 says: " + msg)
}

func main() {
	printerOne := func(msg string) { fmt.Println(msg) }
	callPrinters(
		"Hello",
		printerOne,
		func(msg string) { fmt.Println(msg) },
	)
}
