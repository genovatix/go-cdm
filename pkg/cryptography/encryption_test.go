package cryptography

import (
	"github.com/stretchr/testify/assert"
	"go.dedis.ch/kyber/v3/encrypt/ecies"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/util/random"
	"testing"
)

func TestEciesEncryptionDecryption(t *testing.T) {
	suite := edwards25519.NewBlakeSHA256Ed25519()
	message := []byte("This is a secret message")

	// Generate recipient's keypair
	privateKey := suite.Scalar().Pick(random.New())
	publicKey := suite.Point().Mul(privateKey, nil)

	// Encrypt the message
	encryptedMessage, err := ecies.Encrypt(suite, publicKey, message, suite.Hash)
	assert.NoError(t, err, "Encryption should succeed without errors")

	// Decrypt the message
	decryptedMessage, err := ecies.Decrypt(suite, privateKey, encryptedMessage, suite.Hash)
	assert.NoError(t, err, "Decryption should succeed without errors")

	// Verify the decrypted message matches the original
	assert.Equal(t, message, decryptedMessage, "Decrypted message should match the original")
}
