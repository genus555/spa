# **Simple Password Aggregator**

A simple password aggregator that saves passwords that can be recalled later by their name. Passwords are encrypted before being saved. To access your saved passwords, you will need to provide your username and a temporary password saved on a TOTP Authenticator app of your choice.

## Requirements:
- [Go](https://golang.org/) 1.21+
- TOTP Authenticator App
  * Google Authenticator (iOS/Android)
  * Authy (iOS/Android/Desktop)
  * Microsoft Authenticator (iOS/Android)

## How to use SPA

***!!! IMPORTANT !!!***  
*Make sure to write down/save the encryption key in your .env file elsewhere in case something happens to the original.*
1. Navigate over to the project's root directory in your terminal (e.g. `cd c/Users/your_username/spa`)
2. `go run ./client` from the spa directory.
3. SPA will start and you will be prompted to type in your username.
    1. If you're a new user or there's no user saved to file, you will be prompted to provide a username you're OTP will be saved with.
    2. Then a QR code will be created and displayed at [http://localhost:8080/qr](http://localhost:8080/qr) that can be scanned using a TOTP Authenticator App to get your reusable OTP.
4. You will then be prompted to provide your OTP
5. Finally you will be provided with a list of available commands and how to use them.

### Transferring Information:

*Information on commands `transfer in/out`.*  
**Transfer out:**  
This command is to create a transfer folder containing your .env file which holds the encryption key needed to decrypt your passwords back into readable form.  
*Note:* The transfer folder can also serve as a backup folder  
**Transfer In:**  
Make sure to have your "Transfer" folder inside the spa directory before running this command. This command will copy over the encryption key and database stored inside the transfer folder into the project folder.