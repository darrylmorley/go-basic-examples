package main

import "log"

// The use of pointers in method receivers is common in Go for several reasons:
// It allows the method to modify the fields of the struct it's attached to.
// It's more efficient for large structs as it avoids copying the entire struct each time the method is called.

type myStruct struct {
	FirstName string
}

// Method declaration, received myStruct pointer
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
