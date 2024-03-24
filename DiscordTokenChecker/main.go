package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

var useragents = []string{"Mozilla/5.0 (X11; CrOS x86_64 14541.0.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36", "Mozilla/5.0 (X11; CrOS x86_64 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36", "Mozilla/5.0 (X11; CrOS x86_64 13904.77.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0 Safari/537.36", "Mozilla/5.0 (X11; CrOS aarch64 13099.85.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36", "Mozilla/5.0 (X11; CrOS armv7l 13099.110.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.6167.161 Safari/537.36", "Mozilla/5.0 (X11; CrOS aarch64 13099.85.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0"}

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

	fileGoodTokens, err := os.OpenFile("good_tokens.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileGoodTokens.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := strings.TrimSpace(scanner.Text())
		if token == "" {
			continue
		}
		req := fasthttp.AcquireRequest()
		req.SetRequestURI("https://discord.com/api/v9/users/@me/library")
		req.Header.SetMethod("GET")
		req.Header.Add("Authorization", token)
		req.Header.SetUserAgent(useragents[rand.Intn(len(useragents))])
		resp := fasthttp.AcquireResponse()
		err := fasthttp.Do(req, resp)

		if err != nil {
			fmt.Println("Request failed: ", err)
		}

		statusCode := resp.StatusCode()

		if statusCode == 200 {
			fmt.Println("Good token:", token)
			_, err = fileGoodTokens.WriteString(token + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		} else if statusCode == 401 {
			fmt.Println("Bad token:", token)
		} else if statusCode == 403 {
			fmt.Println("Locked token:", token)
		} else {
			fmt.Println("Unknown error with token:", token)
		}
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}

}
