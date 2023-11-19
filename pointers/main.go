package main

import "log"

// Using a pointer to modify the value of a variable
func main() {
	myString := "Green"

	log.Println("myString is set to", myString)
	// Pass the address of myString to changeUsingPointers
	changeUsingPointers(&myString)
	log.Println("After func call myString is set to", myString)
}

func changeUsingPointers(s *string) {
	// Change the value at the passed memory address to "Red"
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
