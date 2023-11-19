package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// Working with slices
	var mySlice []string
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Adam")
	mySlice = append(mySlice, "Luke")
	sort.Strings(mySlice)
	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	log.Println(numbers[:3])
	log.Println(numbers[2:])

	// Working with maps
	myMap := make(map[string]string)
	myOtherMap := make(map[string]int)
	myUserMap := make(map[string]User)

	myMap["dog"] = "Samson"
	myMap["car"] = "Mercedes"
	log.Println(myMap)
	log.Println(myMap["dog"])

	myOtherMap["First"] = 1
	myOtherMap["Second"] = 2
	log.Println(myOtherMap)
	log.Println(myOtherMap["Second"])

	me := User{
		FirstName: "Darryl",
		LastName:  "Morley",
	}
	myUserMap["me"] = me
	log.Println(myUserMap["me"].FirstName)
	log.Println(myUserMap["me"])
}
