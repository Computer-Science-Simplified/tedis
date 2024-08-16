package main

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/command"
	"github.com/Computer-Science-Simplified/tedis/server/internal/persistence/aol"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 250000; i++ {
		source := rand.NewSource(time.Now().UnixNano())

		r := rand.New(source)

		value := int64(r.Intn(1000000) + 1)

		cmd, err := command.Create("BSTADD", "a", []int64{value})
		if err != nil {
			fmt.Println("could not create command")

			continue
		}

		err = aol.Append(cmd)
		if err != nil {
			fmt.Println("could not run command")
		}
	}
}
