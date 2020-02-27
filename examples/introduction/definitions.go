package main

import (
	"fmt" // <- package
)

type Printer struct{} // <- Type (struct)

type Greeter interface { // <- Type (interface)
	SayHello()
}

func GreetWithGreeter(g Greeter) {
	g.SayHello()
}

func GreetWithFunction(f func()) {
	f()
}

func (p *Printer) printMsg(msg string) { // <- method
	fmt.Println(msg)
}

func (p Printer) SayHello() { // <- method
	p.printMsg("Hello!")
}

func main() { // <- function
	p := Printer{} // <- variable
	GreetWithGreeter(p)
	GreetWithFunction(p.SayHello)
}

func ComplexReceiver(
	f func(string, int, ...map[int]float64),
	enable bool,
) (
	bool, func() error, error,
) {
	// ...
}

type InputFunc func(string, int, ...map[int]float64)

type OutputFunc func() error

func ComplexReceiver(f InputFunc, enable bool) (bool, OutputFunc, error) {

}
