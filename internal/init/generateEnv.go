package init

import (
	"crypto/rand"
)

func createEncKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {return key, err}
	return key, nil
}

func GenerateEncryptionKey() ([]byte, error) {
	byte_key, err := createEncKey()
	if err != nil {return nil, err}
	return byte_key, nil
}