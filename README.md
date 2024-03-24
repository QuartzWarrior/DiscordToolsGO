# Discord Tools - GO

This repository contains Go programs designed to handle and manipulate Discord tokens in GO.

## DiscordTokenFormatter

This program handles and manipulates tokens stored in `tokens.txt` in the format `email:pass:token`. It counts the number of valid tokens and offers conversion to different formats. Run with `go run main.go`.

## DiscordTokenOnliner

This program interacts with the Discord API using tokens from `tokens.txt`. It allows setting the status and game of the Discord user associated with each token. Run with `go run main.go`.

## DiscordTokenChecker

This program validates Discord tokens by making a GET request to the Discord API. Tokens are classified as "good", "bad", or "locked" based on the API response. Good tokens are written to `good_tokens.txt`. Run with `go run main.go`.

## DiscordVanityChecker

This program checks the availability of Discord vanity URLs by making a GET request to the Discord API. Vanities are classified as "available" or "taken" based on the API response. Run with `go run main.go`.

## DiscordNitroChecker

This program validates Discord Nitro gift codes by making a GET request to the Discord API. Codes are classified as "valid" or "invalid" based on the API response. Run with `go run main.go`.

## DiscordPasswordChanger

This Go program changes the password of Discord accounts using account tokens. It makes a PATCH request to the Discord API and confirms the password change based on the API response. Run with `go run main.go`.

## Usage

Ensure all tokens are valid Discord tokens with DiscordTokenChecker. If `tokens.txt` is empty, you will be prompted to input tokens when running the programs. Use the command `go run main.go` to run a program.

## Note

Error handling is included for file operations and API interactions.