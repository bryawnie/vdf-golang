package main

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/chia-network/go-chia-libs/pkg/vdf"
	log "github.com/sirupsen/logrus"
)

const (
	lambda     uint64 = 1024    // The discriminant size
	iterations uint64 = 2000000 // Also denoted as T
	form_size  int    = 100     // The size of the form
)

func getRandomBytes(nBytes int) []byte {
	seed := make([]byte, nBytes)
	rand.Read(seed)
	return seed
}

func serializeInput(input []byte) []byte {
	// Prepend the byte 0x08 to the input
	prependByte := byte(0x08)
	return append([]byte{prependByte}, input...)
}

func main() {
	log.SetLevel(log.DebugLevel)

	log.Info("Sample VDF (Verifiable Delay Function) Implementation in GoLang")

	// Create a random challenge (or seed)
	log.Debug("Creating random challenge")
	challenge := getRandomBytes(64) // Also denoted as Seed

	x := serializeInput(getRandomBytes(form_size - 1)) // Also denoted as Input

	// Evaluate the VDF
	log.Debug("Evaluting VDF")
	outVdf := vdf.Prove(challenge, x, int(lambda), iterations)
	y := outVdf[0:form_size]
	proof := outVdf[form_size:]

	log.Infof("VDF output: %s", hex.EncodeToString(y))
	log.Infof("VDF proof: %s", hex.EncodeToString(proof))

	// Verify the VDF
	log.Debug("Verifying VDF")
	// Create discriminant
	discriminant := vdf.CreateDiscriminant(challenge, int(lambda))
	log.Infof("Discriminant: %s", discriminant)

	recursion := 0 // We do not use recursion for final output verification
	verified := vdf.VerifyNWesolowski(discriminant, x, outVdf, iterations, lambda, uint64(recursion))
	log.Infof("VDF verification: %t", verified)
}
