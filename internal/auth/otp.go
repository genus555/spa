package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/pquerna/otp"
	"github.com/skip2/go-qrcode"
	cl			"github.com/genus555/spa/internal/clientloop"
	database	"github.com/genus555/spa/internal/database"
)

func createKey(username string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:			"spa",
		AccountName:	username,
	})
	if err != nil {return nil, err}
	return key, nil
}

func generateQrCode(key_url string) error {
	http.HandleFunc("/qr", func(w http.ResponseWriter, r *http.Request) {
		png, _ := qrcode.Encode(key_url, qrcode.Medium, 256)
		w.Header().Set("Content-Type", "image/png")
		w.Write(png)
	})

	fmt.Printf("QR code generated\nGo to \"http://localhost:8080/qr\" in your browser to scan it.\n")
	return http.ListenAndServe(":8080", nil)
}

func NewUser(db *database.DB) (string, string, error) {
	username := cl.GetUsername()
	key, err := createKey(username)
	if err != nil {return "", "", err}

	go generateQrCode(key.URL())
	time.Sleep(1 *time.Second)
	return username, key.Secret(), nil
}

func CheckValid(db *database.DB, username, otp_secret string) error {
	for {
		fmt.Printf("Enter %s's passcode:\n", username)
		input := cl.GetPasscode()
		valid, err := Valid(db, username, input, otp_secret)
		if err != nil {return err}
		if !valid {
			fmt.Println("Incorrect passcode")
		} else {
			break
		}
	}

	return nil
}

func Valid(db *database.DB, username, passcode, otp_secret string) (bool, error) {
	return totp.Validate(passcode, otp_secret), nil
}