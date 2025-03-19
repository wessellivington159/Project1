package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42)
	num := rand.Intn(100)
	fmt.Println("Random number:", num)
}
