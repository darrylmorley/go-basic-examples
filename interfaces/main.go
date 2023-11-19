package main

import "fmt"

// An interface is a type that defines a set of methods that can be implemented by other types.
type Animal interface {
	says() string
	legs() int
}

// A struct is a type that groups together zero or more named fields of any type.
type Dog struct {
	Name  string
	Breed string
}

type Monkey struct {
	Name  string
	Color string
}

func main() {
	dog := Dog{
		Name:  "Nim",
		Breed: "Golden Retriever",
	}

	PrintInfo(&dog)

	monkey := Monkey{
		Name:  "Bill",
		Color: "Brown",
	}

	PrintInfo(&monkey)
}

func PrintInfo(a Animal) {
	fmt.Println("this animal says", a.says(), "and has", a.legs(), "legs")
}

// A method is a function that has a receiver, which is the type that the method belongs too.
func (d *Dog) says() string {
	return "Woof"
}

func (d *Dog) legs() int {
	return 4
}

func (m *Monkey) says() string {
	return "Ooh ooh"
}

func (m *Monkey) legs() int {
	return 2
}
