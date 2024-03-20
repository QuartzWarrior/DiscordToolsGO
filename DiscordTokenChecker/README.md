# DiscordTokenChecker

This Go program is designed to validate Discord tokens. It reads tokens from a file named `tokens.txt`, and uses these tokens to make a GET request to the Discord API. The response from the API is used to determine whether each token is valid.

## Features

1. **Token Input**: If the `tokens.txt` file is empty, the program will prompt the user to input tokens. Each token should be entered on a new line.

2. **Token Validation**: The program validates each token by making a GET request to the Discord API. The response from the API is used to determine whether the token is valid.

3. **Token Classification**: The program classifies each token based on the response from the API. Tokens are classified as "good", "bad", or "locked". Good tokens are written to a file named `good_tokens.txt`.

## Error Handling

The program includes error handling for file operations such as opening, writing to, and reading from files. It also handles errors that may occur during the GET request to the Discord API.

## Usage

To run the program, use the command `go run main.go` in the terminal. Follow the prompts to input tokens if the `tokens.txt` file is empty.