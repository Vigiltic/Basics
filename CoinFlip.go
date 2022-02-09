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

	fmt.Println("The coin side turned out to be:", side)
}
