package cryptography

import (
	"fmt"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/encrypt/ecies"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/sign/schnorr"
)

// KeyManager defines the interface for key management, signing, and ZKP.
type KeyManager interface {
	GenerateKeys() (kyber.Scalar, kyber.Point)                           // Generates a private and public key pair
	SignMessage(privateKey kyber.Scalar, message []byte) ([]byte, error) // Signs a message using Schnorr signature
	/*GenerateProof(circuit frontend.API, witness frontend.Circuit) ([]byte, error) // Generates a ZKP
	VerifyProof(circuit frontend.API, proof []byte) (bool, error)    */ // Verifies a ZKP
}

type DefaultKeyManager struct{}

func (km *DefaultKeyManager) GenerateKeys() (kyber.Scalar, kyber.Point) {
	suite := edwards25519.NewBlakeSHA256Ed25519()
	privateKey := suite.Scalar().Pick(suite.RandomStream())
	publicKey := suite.Point().Mul(privateKey, nil)
	return privateKey, publicKey
}

func (km *DefaultKeyManager) SignMessage(privateKey kyber.Scalar, message []byte) ([]byte, error) {
	suite := edwards25519.NewBlakeSHA256Ed25519()
	signature, err := schnorr.Sign(suite, privateKey, message)
	return signature, err
}

func NewKeyManager() KeyManager {
	return &DefaultKeyManager{}
}

var suite = edwards25519.NewBlakeSHA256Ed25519()

func (km *DefaultKeyManager) Encrypt(publicKey kyber.Point, message []byte) ([]byte, error) {

	cipherText, err := ecies.Encrypt(suite, publicKey, message, suite.Hash)
	if err != nil {
		panic(err)
	}
	return cipherText, nil

}

func (km *DefaultKeyManager) Decrypt(privateKey kyber.Scalar, cipherText []byte) ([]byte, error) {
	message, err := ecies.Decrypt(suite, privateKey, cipherText, suite.Hash)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt message: %v", err)
	}
	return message, nil
}
