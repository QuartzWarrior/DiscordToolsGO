package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

func main() {
	file, err := os.OpenFile("codes.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
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
		fmt.Println("Enter your nitro codes (one code per line, use Shift+Enter to go to the next time, press Enter to finish):")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			code := strings.TrimSpace(scanner.Text())
			_, err := file.WriteString(code + "\n")
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

	fileGoodCodes, err := os.OpenFile("good_codes.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileGoodCodes.Close()

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code := strings.TrimSpace(scanner.Text())
		if code == "" {
			continue
		}
		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("GET")
		req.SetRequestURI(fmt.Sprintf("https://discordapp.com/api/v9/entitlements/gift-codes/%s?with_application=false&with_subscription_plan=true", code))
		resp := fasthttp.AcquireResponse()
		fasthttp.Do(req, resp)
		if resp.StatusCode() == 200 {
			fmt.Println("Valid code:", code)
			_, err = fileGoodCodes.WriteString(code + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		} else {
			fmt.Println("Invalid code:", code)
		}
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}

}
