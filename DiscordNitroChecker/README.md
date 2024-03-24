# DiscordNitroChecker

This Go program is designed to validate Discord Nitro gift codes. It reads codes from a file named `codes.txt`, and uses these codes to make a GET request to the Discord API. The response from the API is used to determine whether each code is valid.

## Features

1. **Code Input**: If the `codes.txt` file is empty, the program will prompt the user to input Nitro gift codes. Each code should be entered on a new line.

2. **Code Validation**: The program validates each code by making a GET request to the Discord API. The response from the API is used to determine whether the gift code is valid.

3. **Code Classification**: The program classifies each code based on the response from the API. Codes are classified as "valid" or "invalid". Valid codes are written to a file named `good_codes.txt`.

## Error Handling

The program includes error handling for file operations such as opening, writing to, and reading from files. It also handles errors that may occur during the GET request to the Discord API.

## Usage

To run the program, use the command `go run DiscordNitroChecker.go` in the terminal. Follow the prompts to input codes if the `codes.txt` file is empty.