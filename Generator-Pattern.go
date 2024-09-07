package main

import (
	"fmt"
)

// Generator function that returns a channel of integers (unique IDs)
// and a channel to signal when to stop the generator

func idGenerator(done <-chan struct{}) <-chan int {
	idChan := make(chan int)

	go func() {
		id := 0
		for {
			select {
			case idChan <- id:
				{
					id++
				}
			case <-done:
				{
					close(idChan)
					return;
				}
			}
		}
	}()
	return idChan
}

func generatorTestFunction() {
	// Create the done channel
	done := make(chan struct{})

	// Create the ID generator
	gen := idGenerator(done)

	// Generate and use IDs on demand
	fmt.Println("Requesting IDs...")
	for i := 0; i < 5; i++ {
		id := <-gen // Receive the next ID from the channel
		fmt.Printf("Generated ID: %d\n", id)
	}

	// Simulate more ID requests later
	fmt.Println("Requesting more IDs...")
	for i := 0; i < 3; i++ {
		id := <-gen // Receive the next ID from the channel
		fmt.Printf("Generated ID: %d\n", id)
	}

	// Signal the generator to stop and close the channel
	close(done)

	// Optionally, attempt to read from the channel after closure
	// This will return the zero value for the channel's type (0 in this case)
	// and a boolean indicating the channel is closed (false)
	if id, ok := <-gen; !ok {
		fmt.Println("Generator has stopped, channel is closed")
	} else {
		fmt.Printf("Generated ID after closure: %d\n", id)
	}
}
