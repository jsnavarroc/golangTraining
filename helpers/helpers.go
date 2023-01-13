package helpers

import (
	"math/rand"
	"time"
)

func RandomNumer(n int) int {
	rand.Seed(time.Now().UnixMicro())
	value := rand.Intn(n)
	return value
}