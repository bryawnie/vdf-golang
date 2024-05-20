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
	seed := getRandomSeed()
	length := 100
	discr := vdf.CreateDiscriminant(seed, length)

	fmt.Println("Discriminant: ", discr)
}
