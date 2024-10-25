package main

import (
	"fmt"
	"math/rand"
	"time"
)

var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"سلام دنیا",
	"Привет, мир",
}

func main() {
	// Declare a global variable named helloList of type slice of string

	// Seed random number generator using the current time
	// rand.Seed((time.Now().UnixNano()))
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate a random number in the range of our list
	index := rand.Intn((len(helloList)))

	fmt.Println(helloList[index])
}
