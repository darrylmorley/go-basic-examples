package main

import (
	"basics/channels/helpers"
	"log"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	// Generate a random number
	randomNumber := helpers.RandomNumber(numPool)

	// Pass the random number via the intChan channel
	intChan <- randomNumber
}

func main() {
	// Create a channel of type int - will except only int's
	intChan := make(chan int)

	// Close the channel at end of run
	defer close(intChan)

	// run the CalculateValue func as a go routine - send or receive values via the provided channel (intChan)
	go CalculateValue(intChan)

	// Receive the value from intChan channel
	num := <-intChan

	// log out the value
	log.Println(num)
}
