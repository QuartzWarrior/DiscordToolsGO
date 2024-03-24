package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

type ProvidedData struct {
	Password    string
	NewPassword string
}

type ResponseData struct {
	ID               string   `json:"id"`
	Username         string   `json:"username"`
	Avatar           string   `json:"avatar"`
	Discriminator    string   `json:"discriminator"`
	PublicFlags      int      `json:"public_flags"`
	Types            int      `json:"premium_type"`
	Flags            int      `json:"flags"`
	Banner           *string  `json:"banner"`
	AccentColor      *string  `json:"accent_color"`
	GlobalName       string   `json:"global_name"`
	AvatarDecoration *string  `json:"avatar_decoration_data"`
	BannerColor      *string  `json:"banner_color"`
	MFA              bool     `json:"mfa_enabled"`
	Locale           string   `json:"locale"`
	Email            string   `json:"email"`
	Verified         bool     `json:"verified"`
	Token            string   `json:"token"`
	Phone            string   `json:"phone"`
	NSFW             bool     `json:"nsfw_allowed"`
	Linked           []string `json:"linked_users"`
	Purchased        int      `json:"purchased_flags"`
	Bio              string   `json:"bio"`
	AuthTypes        []string `json:"authenticator_types"`
}

func runner(tokenDataString string, newPassword string, fileFinishedTokens *os.File) {
	tokenData := strings.Split(tokenDataString, ":")

	// email := tokenData[0]
	password := tokenData[1]
	token := tokenData[2]

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://discord.com/api/v9/users/@me")
	req.Header.SetMethod("PATCH")
	req.Header.Set("authority", "discord.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("authorization", token)
	req.Header.SetContentType("application/json")
	req.Header.Set("origin", "https://discord.com")
	req.Header.Set("referer", "https://discord.com/channels/@me")
	req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"122\", \"Not(A:Brand\";v=\"24\", \"Google Chrome\";v=\"122\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Chrome OS\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.SetUserAgent("Mozilla/5.0 (X11; CrOS x86_64 14541.0.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36'")
	req.Header.Set("x-debug-options", "bugReporterEnabled")
	req.Header.Set("x-discord-locale", "en-US")
	req.Header.Set("x-discord-timezone", "America/Denver")
	req.Header.Set("x-super-properties", "eyJvcyI6IiIsImJyb3dzZXIiOiJDaHJvbWUiLCJkZXZpY2UiOiIiLCJzeXN0ZW1fbG9jYWxlIjoiZW4tVVMiLCJicm93c2VyX3VzZXJfYWdlbnQiOiJNb3ppbGxhLzUuMCAoWDExOyBDck9TIHg4Nl82NCAxNDU0MS4wLjApIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMjIuMC4wLjAgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6IjEyMi4wLjAuMCIsIm9zX3ZlcnNpb24iOiIiLCJyZWZlcnJlciI6IiIsInJlZmVycmluZ19kb21haW4iOiIiLCJyZWZlcnJlcl9jdXJyZW50IjoiIiwicmVmZXJyaW5nX2RvbWFpbl9jdXJyZW50IjoiIiwicmVsZWFzZV9jaGFubmVsIjoic3RhYmxlIiwiY2xpZW50X2J1aWxkX251bWJlciI6Mjc0Mzg4LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==")

	reqBody := &ProvidedData{
		Password:    password,
		NewPassword: newPassword,
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error marshalling json", err)
		return
	}

	req.SetBodyRaw(reqBodyBytes)
	resp := fasthttp.AcquireResponse()

	fasthttp.Do(req, resp)

	if resp.StatusCode() == 401 {
		fmt.Println("Invalid token")
		return
	} else {
		var respBody ResponseData
		err = json.Unmarshal(resp.Body(), &respBody)
		if err != nil {
			fmt.Println("Error unmarshaling response: ", err)
			return
		}

		fmt.Println("Success:", respBody.Token)

		_, err := fileFinishedTokens.WriteString(fmt.Sprintf("%s:%s:%s", respBody.Email, newPassword, respBody.Token) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)
}

func main() {
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
		fmt.Println("Enter your tokens (one token per line, in email:password:token format, use Shift+Enter to go to the next time, press Enter to finish):")
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

	fileFinishedTokens, err := os.OpenFile("finished.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileFinishedTokens.Close()

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}

	fmt.Println("Enter the password you'd like to use for every account, or random for a random password.")
	fmt.Print("-> ")
	var newPassword string
	fmt.Scan(&newPassword)
	newPassword = strings.TrimSpace(newPassword)
	if newPassword == "random" {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		length := r.Intn(5) + 8
		bytes := make([]byte, length)
		for i := 0; i < length; i++ {
			bytes[i] = byte(r.Intn(94) + 33)
		}
		newPassword = string(bytes)
	}

	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokenDataString := strings.TrimSpace(scanner.Text())
		if tokenDataString == "" {
			continue
		}

		wg.Add(1)
		go func(tokenDataString string) {
			defer wg.Done()
			runner(tokenDataString, newPassword, fileFinishedTokens)
		}(tokenDataString)
	}

	wg.Wait()
}
