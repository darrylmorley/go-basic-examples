package helpers

import (
	"math/rand"
)

// Generate a random number
func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}
