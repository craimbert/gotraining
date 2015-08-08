// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/G7O-DnJrEA

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

// Add imports.
import (
	"fmt"
	"sync"
)
// Declare a wait group variable.
var wg sync.WaitGroup

// Declare a function for the goroutine. Pass in a name for the
// routine and a channel of type int used to share the value back and forth.
func goroutine(name string, channel chan int ) {
	for {
		// Receive on the channel, waiting for the integer.
		val, ok := <- channel
		// Check if the channel is closed.
		if !ok{
			// Call done on the waitgroup.
			wg.Done()
			// Display the goroutine is finished and return.
			fmt.Printf("Go routine %s finised - channel seems closed\n", name)
			return
		}

		// Display the integer value received.
		fmt.Printf("value received by go routine %s from channel: %d\n", name, val)
		// Check is the value from the channel is 10.
		if val == 10 {
			// Close the channel.
			close(channel)
			// Call done on the waitgroup.
			wg.Done()
			// Display the goroutine is finished and return.
			fmt.Printf("Go routine %s finised - target result achieved\n", name)
			return
		}

		// Increment the value by one and send in back through
		// the channel.
		val++
		channel <- val
	}
}

// main is the entry point for all Go programs.
func main() {
	// Declare and initialize an unbuffered channel of type int.
	channel := make(chan int)
	// Increment the wait group for each goroutine
	// to be created.
	wg.Add(2)
	// Create the two goroutines.
	go goroutine("[ A ]", channel)
	go goroutine("[ B ]", channel)
	// Send the initial integer value into the channel.
	channel <- 1
	// Wait for all the goroutines to finish.
	wg.Wait()
}
