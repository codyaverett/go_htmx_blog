package main

import (
	"fmt"
	"time"
)

// Go does not have colored functions
// e.g. once a function becomes async all other functions that call it won't need to be async
//      like they do in JS/TS

// Producer function that sends integers to a channel
func producer(numbers chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		numbers <- i
		time.Sleep(time.Second) // Simulate some work
	}
	close(numbers) // Close the channel when done producing
}

// Consumer function that receives integers from a channel
func consumer(numbers <-chan int) {
	for num := range numbers {
		fmt.Printf("Consuming: %d\n", num)
		time.Sleep(time.Second * 2) // Simulate some work
	}
}

func main() {
	// Create a channel to communicate between producer and consumer
	numbers := make(chan int)

	// Start the producer and consumer as separate goroutines
	go producer(numbers)
	go consumer(numbers)

	// Wait for some time to allow the goroutines to finish (a quick and dirty way to wait)
	time.Sleep(time.Second * 12)
	fmt.Println("All done!")
}
