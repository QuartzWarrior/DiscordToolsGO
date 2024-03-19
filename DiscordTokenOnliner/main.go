package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

type Auth struct {
	Op int `json:"op"`
	D  struct {
		Token      string `json:"token"`
		Properties struct {
			Os      string `json:"$os"`
			Browser string `json:"$browser"`
			Device  string `json:"$device"`
		} `json:"properties"`
		Presence struct {
			Game   Game   `json:"game"`
			Status string `json:"status"`
			Since  int    `json:"since"`
			Afk    bool   `json:"afk"`
		} `json:"presence"`
	} `json:"d"`
	S interface{} `json:"s"`
	T interface{} `json:"t"`
}

type Game struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	Url  string `json:"url,omitempty"`
}

type Hello struct {
	Op int `json:"op"`
	D  struct {
		HeartbeatInterval int `json:"heartbeat_interval"`
	} `json:"d"`
}

func online(token, game, typeStr, status string, wg *sync.WaitGroup) {
	defer wg.Done()
	c, _, err := websocket.DefaultDialer.Dial("wss://gateway.discord.gg/?v=6&encoding=json", nil)

	if err != nil {
		fmt.Println("Error dialing websocket: ", err)
		return
	}

	_, message, err := c.ReadMessage()
	if err != nil {
		fmt.Println("Error reading Hello message: ", err)
		return
	}

	var hello Hello
	err = json.Unmarshal(message, &hello)
	if err != nil {
		fmt.Println("Error unmarshaling Hello message: ", err)
		return
	}

	heartbeatInterval := time.Duration(hello.D.HeartbeatInterval) * time.Millisecond

	var gamejson Game

	if status == "random" {
		statuses := []string{"online", "dnd", "idle"}
		status = statuses[rand.Intn(len(statuses))]
	}

	switch typeStr {
	case "1":
		gamejson = Game{Name: game, Type: 0}
	case "2":
		gamejson = Game{Name: game, Type: 1, Url: game}
	case "3":
		gamejson = Game{Name: game, Type: 2}
	case "4":
		gamejson = Game{Name: game, Type: 3}
	}

	auth := Auth{
		Op: 2,
		D: struct {
			Token      string `json:"token"`
			Properties struct {
				Os      string `json:"$os"`
				Browser string `json:"$browser"`
				Device  string `json:"$device"`
			} `json:"properties"`
			Presence struct {
				Game   Game   `json:"game"`
				Status string `json:"status"`
				Since  int    `json:"since"`
				Afk    bool   `json:"afk"`
			} `json:"presence"`
		}{
			Token: token,
			Properties: struct {
				Os      string `json:"$os"`
				Browser string `json:"$browser"`
				Device  string `json:"$device"`
			}{
				Os:      "linux",
				Browser: "RTB",
				Device:  "linux Device",
			},
			Presence: struct {
				Game   Game   `json:"game"`
				Status string `json:"status"`
				Since  int    `json:"since"`
				Afk    bool   `json:"afk"`
			}{
				Game:   gamejson,
				Status: status,
				Since:  0,
				Afk:    false,
			},
		},
		S: nil,
		T: nil,
	}

	authJson, err := json.Marshal(auth)
	if err != nil {
		fmt.Println("Error marshaling auth JSON:", err)
		return
	}
	err = c.WriteMessage(websocket.TextMessage, authJson)
	if err != nil {
		fmt.Println("Error writing auth message:", err)
		return
	}

	ack := map[string]interface{}{"op": 1, "d": nil}
	ackJson, err := json.Marshal(ack)
	if err != nil {
		fmt.Println("Error marshaling ack JSON:", err)
		return
	}

	for {
		time.Sleep(heartbeatInterval)
		if typeStr == "random" {
			statuses := []string{"online", "dnd", "idle"}
			auth.D.Presence.Status = statuses[rand.Intn(len(statuses))]
			authJson, err = json.Marshal(auth)
			if err != nil {
				fmt.Println("Error marshalling auth data", err)
				return
			}
			err = c.WriteMessage(websocket.TextMessage, authJson)
			if err != nil {
				fmt.Println("Error writing auth message:", err)
				break
			}
		}
		err = c.WriteMessage(websocket.TextMessage, ackJson)
		if err != nil {
			fmt.Println("Error writing ack message:", err)
			break
		}
		fmt.Println("Sent keepalive message")
	}
	defer c.Close()
}

func main() {
	fmt.Println("Options: [1] Playing | [2] Streaming | [3] Watching | [4] Listening\nYour Choice > ")
	var typeStr string
	fmt.Scanln(&typeStr)
	typeStr = strings.TrimSpace(typeStr)

	var game string
	if typeStr == "2" {
		fmt.Println("Info: Type what you want the twitch stream to be. Must start with https://twitch.tv\nTwitch > ")
		fmt.Scanln(&game)
		game = strings.TrimSpace(game)
	} else {
		fmt.Println("Info: Type what you want the game to be\nGame > ")
		fmt.Scanln(&game)
		game = strings.TrimSpace(game)
	}

	fmt.Println("Info: Type what you want the status to be (online, dnd, idle, random)\nStatus > ")
	var status string
	fmt.Scanln(&status)
	status = strings.TrimSpace(status)

	file, err := os.OpenFile("tokens.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	if fileInfo.Size() == 0 {
		fmt.Println("Enter your tokens (one token per line, use Shift+Enter to go to the next time, press Enter to finish):")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			token := strings.TrimSpace(scanner.Text())
			_, err := file.WriteString(token + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := strings.TrimSpace(scanner.Text())
		wg.Add(1)
		go online(token, game, typeStr, status, &wg)
	}

	wg.Wait()

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		os.Exit(0)
	}()

	for {
		fmt.Println("[!] Tokens Online")
		time.Sleep(time.Minute)
	}
}
