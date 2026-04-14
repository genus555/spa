# **Simple Password Aggregator**

A simple password aggregator that saves passwords that can be recalled later by their name. Passwords are encrypted before being saved. To access your saved passwords, you will need to provide your username and a temporary password saved on a TOTP Authenticator app of your choice.  
  
## Motivation  
Everything on the internet requires an account and best security practice is to use a different password for every website/account. For me personally, I like playing games on multiple accounts because I like swapping between differnt IGN's, meaning I have several passwords linked to multiple differnt usernames I need to remember. It's very easy to forget your password or which password is associated with which account so I made SPA to help save passwords with a name to remember which account/website they're associated with.

## Requirements:
- [Go](https://golang.org/) 1.21+
- TOTP Authenticator App
  * Google Authenticator (iOS/Android)
  * Authy (iOS/Android/Desktop)
  * Microsoft Authenticator (iOS/Android)

## Quick Start  
1. Navigate over to the project's root directory in your terminal (e.g. `cd c/Users/your_username/spa`)
2. `go run ./client` from the spa directory.
3. SPA will start and you will be prompted to type in your username.
    1. If you're a new user or there's no user saved to file, you will be prompted to provide a username you're OTP will be saved with.
    2. Then a QR code will be created and displayed at [http://localhost:8080/qr](http://localhost:8080/qr) that can be scanned using a TOTP Authenticator App to get your reusable OTP.
4. You will then be prompted to provide your OTP
5. Finally you will be provided with a list of available commands and how to use them.  
  
## Usage:  
***!!! IMPORTANT !!!***  
*Make sure to write down/save the encryption key in your .env file elsewhere in case something happens to the original.*  
### Register:  
*Usage: `register [password_name] [password]`*  
Registers a new password to the passwords database.  
### Get:  
*Usage: `get [password_name]`*  
Retrieves password associated with the given password name.  
### Delete:  
*Usage: `delete [password_name]`*  
Deletes the password associated with the given password name.  
### List:  
*Usage: `list`*  
Lists all password names saved onto the passwords database.  
### Transfer in/out:  
*Usage: `transfer [in/out]`.*  
**Transfer out:**  
Creates a transfer folder containing your database file which holds your user information, passwords, and encryption key.  
*Note:* The transfer folder can also serve as a backup folder  
**Transfer In:**  
Make sure to have your "Transfer" folder inside the spa directory before running this command. This command will copy over the database file inside the transfer folder into the project folder.  
### Deleteuser:  
*Usage: `deleteuser [username] [passcode]`*  
After verifying the username and passcode provided are correct, this command will delete the user on file and stop the program. To create a new user, simply restart the program and you will be prompted to provide a new username.  
*Note:* Deleteuser is to change the user on file and does not wipe the passwords saved to database. To delete saved passwords please use the `delete` command instead.  
### Help:  
*Usage: `help`*  
Provides a list of available commands and how to use them.  
### Quit:  
*Usage: `quit`*  
Stops the program  

## Contributing  
### Submit a pull request  
If you'd like to contribute, fork the repository and open a pull request to the `main` branch.