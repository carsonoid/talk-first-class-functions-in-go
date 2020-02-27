package main

import "fmt"

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

func main() {
	printerMap := map[string]PrintFunc{
		"printerOne": func(msg string) { fmt.Println(msg) },
		"printerTwo": func(msg string) { fmt.Println(msg) },
	}

	printerMap["printerOne"]("Hello")
	printerMap["printerTwo"]("Howdy")

	printerMap["lazyPrinter"] = printerMap["printerOne"]
	printerMap["lazyPrinter"]("Hi")

	fmt.Println("\nRoll call:")
	for name, print := range printerMap {
		print(name + " is standing by")
	}
}
