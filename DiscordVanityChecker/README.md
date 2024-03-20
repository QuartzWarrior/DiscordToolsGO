# DiscordVanityChecker

This Go program is designed to check the availability of Discord vanity URLs. It reads vanities from a file named `list.txt`, and uses these vanities to make a GET request to the Discord API. The response from the API is used to determine whether each vanity URL is available.

## Features

1. **Vanity Input**: If the `list.txt` file is empty, the program will prompt the user to input vanities. Each vanity should be entered on a new line.

2. **Vanity Check**: The program checks the availability of each vanity by making a GET request to the Discord API. The response from the API is used to determine whether the vanity URL is available.

3. **Vanity Classification**: The program classifies each vanity based on the response from the API. Vanities are classified as "available" or "taken". Available vanities are written to a file named `available.txt`, and taken vanities are written to `unavailable.txt`.

## Error Handling

The program includes error handling for file operations such as opening, writing to, and reading from files. It also handles errors that may occur during the GET request to the Discord API.

## Usage

To run the program, use the command `go run main.go` in the terminal. Follow the prompts to input vanities if the `list.txt` file is empty.