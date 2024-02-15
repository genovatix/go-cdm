package identity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddressSerializationDeserialization(t *testing.T) {
	// Example key for encryption/decryption (if applicable)
	//	var key []byte // Assume this is properly initialized for your encryption scheme

	// Initialize test data for RawAddress
	chainCode := make([]byte, ChainCodeLength)
	// Example chain code, public key, private key, and seed values
	// These should be randomly generated or securely created in actual usage
	for i := range chainCode {
		chainCode[i] = byte(i % 256)
	}
	pubKey := make([]byte, PubKeyLength)
	privKey := make([]byte, PrivKeyLength)
	timestamp := time.Now().Unix()
	seed := NewSeed(SeedLength)

	// Initialize RawAddress with test data
	rawAddr := RawAddress{
		chainCode: chainCode,
		ver:       1,
		timestamp: timestamp,
		pubKey:    pubKey,
		privKey:   privKey,
	}
	copy(rawAddr.seed[:], seed)

	// Convert RawAddress to Address (serialize)
	address := Address{
		raw: rawAddr.ToBytes(),
	}

	// Optionally encrypt address here if your design requires it

	// Convert Address back to RawAddress (deserialize)
	deserializedRawAddr, err := address.Raw()
	assert.NoError(t, err, "Deserialization should not fail")

	// Assertions to verify the integrity of the data through the process
	assert.Equal(t, rawAddr.chainCode, deserializedRawAddr.(*RawAddress).chainCode, "Chain codes do not match")
	assert.Equal(t, rawAddr.ver, deserializedRawAddr.(*RawAddress).ver, "Version numbers do not match")
	assert.Equal(t, rawAddr.timestamp, deserializedRawAddr.(*RawAddress).timestamp, "Timestamps do not match")
	assert.Equal(t, rawAddr.pubKey, deserializedRawAddr.(*RawAddress).pubKey, "Public keys do not match")
	assert.Equal(t, rawAddr.privKey, deserializedRawAddr.(*RawAddress).privKey, "Private keys do not match")
	assert.Equal(t, rawAddr.seed, deserializedRawAddr.(*RawAddress).seed, "Seeds do not match")
}
