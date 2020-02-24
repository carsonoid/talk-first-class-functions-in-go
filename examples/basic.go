package main

import (
	"fmt"
)

func getMessage() string {
	return "Hello"
}

func getPrinter() func(string) {
	p := func(msg string) {
		fmt.Println(msg)
	}
	return p
}

func printWith(p func(string), msg string) {
	p(msg)
}

func main() {
	m := getMessage() // m == string
	p := getPrinter() // p == func(string)
	printWith(p, m)
}
