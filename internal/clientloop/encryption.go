package clientloop

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"bytes"
	"io"
)

func TestEncryptPW(key []byte, inputs []string) error {
	encrypted, err := encryptPW(key, inputs[1])
	if err != nil {return err}

	fmt.Println(encrypted)

	decrypted, err := decryptPW(key, encrypted)
	if err != nil {return err}
	fmt.Println(decrypted)
	return nil
}

func encryptPW(key []byte, pw string) ([]byte, error) {
	fmt.Printf("Unencrypted PW: %v\n", pw)
	byte_pw := []byte(pw)

	block, err := aes.NewCipher(key)
	if err != nil {return nil, err}

	padding := aes.BlockSize - len(byte_pw)%aes.BlockSize
	paddedtext := append(byte_pw, bytes.Repeat([]byte{byte(padding)}, padding)...)

	ciphertext := make([]byte, aes.BlockSize+len(paddedtext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {return nil, err}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedtext)

	return ciphertext, nil
}

func decryptPW(key []byte, enc_pw []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {return "", err}

	if len(enc_pw) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := enc_pw[:aes.BlockSize]
	enc_pw = enc_pw[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(enc_pw, enc_pw)

	pad := int(enc_pw[len(enc_pw)-1])
	
	pw := enc_pw[:len(enc_pw)-pad]
	str_pw := string(pw)
	return str_pw, nil
}