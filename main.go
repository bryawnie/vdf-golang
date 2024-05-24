package main

import (
	"crypto/rand"
	"fmt"

	"github.com/chia-network/go-chia-libs/pkg/vdf"
)

func getRandomSeed() []byte {
	seed := make([]byte, 16)
	rand.Read(seed)
	return seed
}

func main() {
	// Create a new VDF instance
	fmt.Println("Starting")
	seed := getRandomSeed()
	length := 100
	fmt.Println("Creating discriminant")
	discr := vdf.CreateDiscriminant(seed, length)
	fmt.Println("Discriminant: ", discr)
}
