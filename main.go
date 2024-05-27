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

var xHex = "77254bc02fe2e06fc073dd7a1ea6f55be1b7499211b4af85308b0ed4145c3d403c5753d8d4d93cd342b68a1e369fe2dba82de2f2a9610b78a4576f4ffafbff0a"
var chalHex = "55b7b686870a02be75cc089cbe20f464"

func getRandomBytes(nBytes int) []byte {
	seed := make([]byte, nBytes)
	rand.Read(seed)
	return seed
}

func serializeInput(input []byte) []byte {
	if len(input) >= form_size-1 {
		log.Fatalf("Input is too large, must be less than %d bytes", form_size-1)
		log.Exit(1)
	}
	expandedInput := append(make([]byte, form_size-len(input)-1), input...)
	// Prepend the byte 0x08 to the input (VDF requires this byte to be present at the beginning of the input)
	prependByte := byte(0x08)
	return append([]byte{prependByte}, expandedInput...)
}

func EvalVDF(challenge []byte, x []byte) ([]byte, []byte) {
	outVdf := vdf.Prove(challenge, x, int(lambda), iterations)
	y := outVdf[0:form_size]
	proof := outVdf[form_size:]
	return y, proof
}

func VerifyVDF(challenge []byte, x []byte, y []byte, proof []byte) bool {
	// Create discriminant
	discriminant := vdf.CreateDiscriminant(challenge, int(lambda))
	// Verify the VDF
	recursion := 0 // We do not use recursion for final output verification
	return vdf.VerifyNWesolowski(discriminant, x, append(y, proof...), iterations, lambda, uint64(recursion))
}

func runExample(useFixedInput bool) {
	var xRaw, challenge []byte

	log.Debug("Creating random challenge") // Also denoted as Hex
	if useFixedInput {
		challenge, _ = hex.DecodeString(chalHex)
	} else {
		challenge = getRandomBytes(64)
	}
	log.Info("Challenge: ", hex.EncodeToString(challenge))

	log.Debug("Assigning input X")
	if useFixedInput {
		xRaw, _ = hex.DecodeString(xHex)
	} else {
		xRaw = getRandomBytes(form_size - 1)
	}
	x := serializeInput(xRaw)
	log.Info("Input: ", hex.EncodeToString(x))

	// Evaluate the VDF
	log.Debug("Evaluating VDF")
	y, proof := EvalVDF(challenge, x)

	log.Infof("VDF output: %s", hex.EncodeToString(y))
	log.Infof("VDF proof: %s", hex.EncodeToString(proof))

	// Verify the VDF
	log.Debug("Verifying VDF")
	verified := VerifyVDF(challenge, x, y, proof)
	log.Infof("VDF verification: %t", verified)
}

func main() {
	log.SetLevel(log.DebugLevel)

	// Run a random example
	runExample(true)
}
