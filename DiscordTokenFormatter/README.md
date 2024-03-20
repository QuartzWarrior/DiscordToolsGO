# DiscordTokenFormatter

This Go program is designed to handle and manipulate tokens that are stored in a specific format. The tokens are stored in a file named `tokens.txt` and the format is `email:pass:token`.

## Features

1. **Token Input**: If the `tokens.txt` file is empty, the program will prompt the user to input tokens in the `email:pass:token` format. Each token should be entered on a new line.

2. **Token Count**: The program counts the number of valid tokens (non-empty lines) in the `tokens.txt` file.

3. **Token Conversion**: The program offers the user three options for converting the format of the tokens. The user can choose to convert the tokens to:
    - `token` format
    - `token:pass` format
    - `token:email:pass` format

The converted tokens are written to an `output.txt` file.

## Error Handling

The program includes error handling for file operations such as opening, writing, and reading files. It also handles errors that may occur during user input and token conversion.

## Usage

To run the program, use the command `go run main.go` in the terminal. Follow the prompts to input tokens or convert existing tokens.

## Note

All tokens must be in `email:pass:token` format.