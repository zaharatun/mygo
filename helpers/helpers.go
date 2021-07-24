package helpers

import "math/rand"

func RendomNumber(n int) int {
	value := rand.Intn(n)
	return value
}