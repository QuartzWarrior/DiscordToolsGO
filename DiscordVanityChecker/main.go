package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

func main() {
	file, err := os.OpenFile("list.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
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
		fmt.Println("Enter your vanities (one vanity per line, use Shift+Enter to go to the next time, press Enter to finish):")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			vanity := strings.TrimSpace(scanner.Text())
			_, err := file.WriteString(vanity + "\n")
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

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}

	availableFile, err := os.OpenFile("available.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return
	}
	defer availableFile.Close()

	unavailableFile, err := os.OpenFile("unavailable.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return
	}
	defer unavailableFile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vanity := strings.TrimSpace(scanner.Text())
		if vanity == "" {
			continue
		}
		req := fasthttp.AcquireRequest()
		req.SetRequestURI(fmt.Sprintf("https://discord.com/api/v9/invites/%s", vanity))
		req.Header.SetMethod("GET")
		resp := fasthttp.AcquireResponse()
		fasthttp.Do(req, resp)

		statusCode := resp.StatusCode()

		if statusCode == 200 {
			fmt.Printf("Vanity is taken: %s\n", vanity)
			_, err = unavailableFile.WriteString(vanity + "\n")
		} else {
			fmt.Printf("Vanity is available: %s\n", vanity)
			_, err = availableFile.WriteString(vanity + "\n")
		}
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}
	}
}
