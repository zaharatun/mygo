package helpers

import "math/rand"
import "time"

func RendomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}