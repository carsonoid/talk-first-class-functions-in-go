package main

import (
	"fmt"
	"strings"
)

// ModifyFunc is any function or method that takes a string and returns a string
type ModifyFunc func(string) string

func main() {
	modifyFuncs := []ModifyFunc{
		func(msg string) string { return msg + " is cool" },
		func(msg string) string { return strings.Title(msg) },
	}

	exclaimString := func(msg string) string {
		return msg + "!"
	}

	modifyFuncs = append(modifyFuncs, exclaimString)

	name := "carson"
	for _, modify := range modifyFuncs {
		name = modify(name)
	}
	fmt.Println(name)
}
