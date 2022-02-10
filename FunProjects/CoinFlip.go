package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	coin := []string{
		"heads",
		"tails",
	}

	rand.Seed(time.Now().UnixNano())

	// flip the coin
	side := coin[rand.Intn(len(coin))]

	fmt.Printf("side: %s\n", side)
	//Pretend this works in the playground
}
