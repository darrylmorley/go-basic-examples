package main

import "log"

func main() {
	isTrue := true

	// if else example
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	// switch example
	myVar := "horse"
	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")
	case "dog":
		log.Println("myVar is set to dog")
	case "fish":
		log.Println("myVar is set to fish")
	default:
		log.Println("myVar is something else")
	}
}
