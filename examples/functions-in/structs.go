package main

import "fmt"

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

type MultiPrinter struct {
	printerOne  PrintFunc
	printerTwo  PrintFunc
	lazyPrinter PrintFunc
}

func main() {
	mp := MultiPrinter{
		printerOne: func(msg string) { fmt.Println(msg) },
		printerTwo: func(msg string) { fmt.Println(msg) },
	}

	mp.printerOne("Hello!")
	mp.printerTwo("Howdy!")

	mp.lazyPrinter = mp.printerTwo
	mp.lazyPrinter("Hi!")
}
