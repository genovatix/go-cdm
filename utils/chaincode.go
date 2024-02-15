package utils

import (
	"crypto/rand"
	"fmt"
)

const (
	// Define sizes for each component of RawAddress
	VerSize         = 4  // Size in bytes for the version field (uint32)
	TimestampSize   = 8  // Size in bytes for the timestamp field (int64)
	ChainCodeLength = 32 // Example length of the chain code in bytes
	PubKeyLength    = 64 // Example length of the public key in bytes, adjust based on your key format
	PrivKeyLength   = 64 // Example length of the private key in bytes, adjust based on your key format
	SeedLength      = 32 // Length of the seed in bytes

	// ExpectedTotalLength Calculate the expected total length of RawAddress when serialized
	ExpectedTotalLength = VerSize + TimestampSize + ChainCodeLength + PubKeyLength + PrivKeyLength + SeedLength
)

func GetChainCode() []byte {
	chainCode := make([]byte, ChainCodeLength)
	_, err := rand.Read(chainCode)
	if err != nil {
		fmt.Println("Error generating chain code:", err)
		return nil
	}
	return chainCode
}
