# DiscordPasswordChanger

This Go program is designed to change the password of Discord accounts. It reads account tokens from a file named `tokens.txt`, and uses these tokens to make a PATCH request to the Discord API. The response from the API is used to determine whether the password change was successful.

## Features

1. **Token Input**: If the `tokens.txt` file is empty, the program will prompt the user to input account tokens. Each token should be entered on a new line.

2. **Password Change**: The program changes the password of each account by making a PATCH request to the Discord API. The user can specify a new password for all accounts, or choose to assign a random password to each account.

3. **Password Change Confirmation**: The program confirms each password change by examining the response from the API. If the password change is successful, the new token is written to a file named `finished.txt`.

## Error Handling

The program includes error handling for file operations such as opening, writing to, and reading from files. It also handles errors that may occur during the PATCH request to the Discord API.

## Usage

To run the program, use the command `go run main.go` in the terminal. Follow the prompts to input tokens if the `tokens.txt` file is empty, and to specify a new password.

## Note

All tokens must be valid Discord account tokens.