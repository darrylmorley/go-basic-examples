package main

import "log"

func main() {
	// Simple for loop example
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	// Loop through a slice example
	animals := []string{"dog", "fish", "horse", "cow", "chicken", "cat"}

	for _, animal := range animals {
		log.Println(animal)
	}

	// Loop through a map example
	cars := make(map[int]string)
	cars[1] = "bmw"
	cars[2] = "ford"
	cars[3] = "audi"

	for _, car := range cars {
		log.Println(car)
	}

	// Loop through a string [byte slice] example
	firstLine := "Once upon a time"
	for _, l := range firstLine {
		log.Println(l)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	var users []User
	users = append(users, User{"John", "Smith", "jsmith@gmail.com", 21})
	users = append(users, User{"Mary", "Jones", "mjones@gmail.com", 23})
	users = append(users, User{"Mark", "Smith", "jsmith@gmail.com", 20})

	for _, u := range users {
		log.Println(u.FirstName, u.Email)
	}
}
