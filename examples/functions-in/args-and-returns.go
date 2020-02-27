package main

import "fmt"

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

func combinePrinters(p1, p2 PrintFunc) PrintFunc {
	return func(msg string) {
		p1("p1 says: " + msg)
		p2("p2 says: " + msg)
	}
}

func combineAllPrinters(printers ...PrintFunc) PrintFunc {
	return func(msg string) {
		for i, p := range printers {
			p(fmt.Sprintf("printer at %d says %s", i, msg))
		}
	}
}

func main() {
	printerOne := func(msg string) { fmt.Println(msg) }

	bothSay := combinePrinters(
		printerOne,
		func(msg string) { fmt.Println(msg) },
	)

	bothSay("Hello")

	everyoneSay := combineAllPrinters(
		func(msg string) { fmt.Println("a: " + msg) },
		func(msg string) { fmt.Println("b: " + msg) },
		func(msg string) { fmt.Println("c: " + msg) },
		func(msg string) { fmt.Println("d: " + msg) },
		func(msg string) { fmt.Println("e: " + msg) },
	)

	everyoneSay("Hello")
}
