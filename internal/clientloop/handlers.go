package clientloop

import (
	"fmt"
)

func HandleRegister(key []byte, inputs []string) error {
	if len(inputs) != 3 {
		fmt.Println("Too many/few arguments")
		fmt.Println("Register Usage: rgister [password_name] [password]")
		return nil
	}

	pw_id := inputs[1]
	pw := inputs[2]
	//test print line
	fmt.Printf("pw_id: %s\npw: %s\n", pw_id, pw)

	enc_pw, err := encryptPW(key, pw)
	if err != nil {return err}

	//add password to sql here
	return nil
}