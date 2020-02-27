package main

import "fmt"

type Greeter interface {
	SayHello()
}

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

// PrintFunc implements the Greeter interface
func (pf PrintFunc) SayHello() {
	if pf != nil {
		pf("Hello")
	} else {
		fmt.Println("Hello from the default printer")
	}
}

func getPrinterGreeter() PrintFunc {
	return func(msg string) { fmt.Println(msg, "from a custom printer") }
}

func main() {
	p := new(PrintFunc)
	p.SayHello()

	p2 := getPrinterGreeter()
	p2.SayHello()

	p3 := PrintFunc(func(msg string) { fmt.Println(msg, "from a forcibly-typed printer") })
	p3.SayHello()
}
