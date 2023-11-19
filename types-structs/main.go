package main

import (
	"log"
	"time"
)

// Example using a struct, also demonstrates that undefined values are zero values in Go
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthdate   time.Time
}

func main() {
	user := User{
		FirstName: "Dave",
		LastName:  "Davis",
	}

	log.Println(user.FirstName)
	log.Println(user.LastName)
	log.Println(user.Age)
	log.Println(user.Birthdate)
}
