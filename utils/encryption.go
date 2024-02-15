package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type Encapsulation interface {
	Encrypt(data []byte, key []byte) ([]byte, error)
	Decrypt(data []byte, key []byte) ([]byte, error)
}

func Encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encrypted := gcm.Seal(nonce, nonce, data, nil)
	return encrypted, nil
}

func Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(data) < gcm.NonceSize() {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

type Encapsulatable interface {
	ToBytes() []byte
}

type Decapsulatable interface {
	FromBytes(data []byte) (interface{}, error)
	Bytes() []byte
	Decapsulated() interface{}
}

func Encapsulate(source Encapsulatable, key []byte) (dst []byte, err error) {

	rawData := source.ToBytes()
	dst, err = Encrypt(rawData, key)
	if err != nil {
		return nil, err
	}
	return

}

func Decapsulate(source Decapsulatable, key []byte) (interface{}, error) {
	decryptedData, err := Decrypt(source.Bytes(), key)
	if err != nil {
		return nil, err
	}

	dst, err := source.FromBytes(decryptedData)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize: %v", err)
	}

	return dst, err
}
