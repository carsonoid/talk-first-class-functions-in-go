package main

import (
	"fmt"
)

// PrintFunc is any function or method that takes a string
type PrintFunc func(string)

func main() {
	// Create a channel to signal that we are done
	done := make(chan struct{})

	// Make a channel that can buffer one print func at a time
	printerChan := make(chan PrintFunc, 1)

	// Create an anonymous function and start it as a dedicated goroutine
	go func() {
		// Loop forever
		for {
			// Watch all cases, block until one returns
			select {
			case <-done: // done signal means break out of the loop
				break
			case p := <-printerChan: // when given a printer via the printerChan, say hello with it
				p("Hello!")
			}
		}
	}()

	// Create a variable for the "printerOne" function
	printerOne := func(msg string) { fmt.Println("printerOne says:", msg) }

	// Send it over the channel
	printerChan <- printerOne

	// Send another printer over the channel without assigning it
	printerChan <- func(msg string) { fmt.Println("printerTwo says:", msg) }

	// create and send lots of functions over the channel
	for i := 0; i < 10; i++ {
		printerChan <- func(msg string) {
			fmt.Println(
				fmt.Sprintf("function %d says", i),
				msg,
			)
		}
	}

	// Tell the background process we are done.
	done <- struct{}{}
}
