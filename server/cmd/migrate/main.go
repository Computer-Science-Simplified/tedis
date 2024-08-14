package main

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/persistence/aol"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 1000000; i++ {
		source := rand.NewSource(time.Now().UnixNano())

		r := rand.New(source)

		value := int64(r.Intn(1000000) + 1)

		aol.Write("bstadd", "a", []int64{value})
	}
}
