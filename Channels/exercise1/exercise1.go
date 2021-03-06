// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"sync"
)

// Add imports.

func main() {

	// Create an unbuffered channel.
	share := make(chan int)

	// Create the WaitGroup and add a count
	// of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	// Launch the goroutine and handle Done.
	go func() {
		goroutine("madhav", share)
		wg.Done()
	}()

	// Launch the goroutine and handle Done.

	go func() {
		goroutine("Keshav", share)
		wg.Done()
	}()

	// Send a value to start the counting.
	share <- 1

	// Wait for the program to finish.
	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(name string, share chan int) {
	for {

		// Wait for the value to be sent.
		// If the channel was closed, return.
		value, ok := <-share
		if !ok {

			// If the channel was closed, return.
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}
		// Display the value.
		fmt.Printf("Goroutine %s Inc %d\n", name, value)

		// Terminate when the value is 10.
		if value == 10 {
			close(share)
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}

		// Increment the value and send it
		// over the channel.
		share <- (value + 1)
	}
}
