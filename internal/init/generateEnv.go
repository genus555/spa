package init

import (
	"os"
	"crypto/rand"
	"encoding/base64"
)

func createEnvKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {return key, err}
	return key, nil
}

func writeToEnv(key []byte) error {
	str_key := base64.StdEncoding.EncodeToString(key)
	keyname := "ENCRYPTION_KEY"
	os.Setenv(keyname, str_key)

	f, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {return err}
	defer f.Close()

	_, err = f.WriteString(keyname + "=" + str_key + "\n")
	if err != nil {return err}
	return nil
}

func GenerateEnv() error {
	exist := os.Getenv("ENCRYPTION_KEY")
	if exist == "" {
		byte_key, err := createEnvKey()
		if err != nil {return err}
		err = writeToEnv(byte_key)
		if err != nil {return err}
		return nil
	} else {
		return nil
	}
}

func DecodeKey(key string) ([]byte, error) {
	byte_key, err := base64.StdEncoding.DecodeString(key)
	if err != nil {return byte_key, err}
	return byte_key, nil
}