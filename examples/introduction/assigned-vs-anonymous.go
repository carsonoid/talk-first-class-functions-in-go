package main

import (
	"fmt"
)

func printWith(p func(string), msg string) {
	p(msg)
}

func main() {
	pfunc := func(msg string) { fmt.Println(msg) }
	printWith(pfunc, "Printed with an anonymous function assigned to a variable")

	printWith(
		func(msg string) {
			fmt.Println(msg)
		},
		"Printed with an anonymous function without assignment",
	)
}
