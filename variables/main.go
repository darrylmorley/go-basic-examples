package main

import "fmt"

// Example of using variables in Go
func main() {
	fmt.Println("Hello World.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye!"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println(whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
