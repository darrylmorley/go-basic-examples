package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100, 20)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("result of division is", result)
}

func divide(x, y float64) (float64, error) {
	var result float64

	if y == 0 {
		return result, errors.New("can not divide by 0")
	}

	result = x / y
	return result, nil
}
