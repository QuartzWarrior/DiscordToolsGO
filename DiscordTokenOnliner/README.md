# DiscordTokenOnliner

This Go program is designed to interact with the Discord API using websockets. It reads tokens from a file named `tokens.txt`, and uses these tokens to authenticate with the Discord API and set the status of the corresponding user.

## Features

1. **Token Input**: If the `tokens.txt` file is empty, the program will prompt the user to input tokens. Each token should be entered on a new line.

2. **Status Setting**: The program allows the user to set the status of the Discord user associated with each token. The status can be set to "online", "dnd" (do not disturb), "idle", or "random". If "random" is chosen, the status will be randomly chosen from the other three options at regular intervals.

3. **Game Setting**: The program allows the user to set the game that the Discord user is playing. The game can be set to "Playing", "Streaming", "Watching", or "Listening". If "Streaming" is chosen, the user is prompted to enter a Twitch stream URL.

## Error Handling

The program includes error handling for file operations such as opening and writing to files. It also handles errors that may occur during websocket operations, such as dialing the websocket, reading messages, and writing messages.

## Usage

To run the program, use the command `go run main.go` in the terminal. Follow the prompts to input tokens, choose a status, and choose a game.

## Note

All tokens must be valid Discord tokens.