package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// generate a random number
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(100)

	fmt.Println("Random number", randomNumber)
	if randomNumber < 33 {
		fmt.Println("small number (number is < 33)")
	} else if randomNumber < 66 {
		fmt.Println("medium number (number is < 66)")
	} else {
		fmt.Println("large number (number is >=66)")
	}

	randomRange := rand.Intn(5)
	fmt.Println("Random range", randomRange)

	i := 0
	for i < randomRange {
		fmt.Println("number ", i)
		i++
	}
}
