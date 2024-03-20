package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	count := 0
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
		fmt.Println("Enter your tokens in email:pass:token format, (one token per line, use Shift+Enter to go to the next time, press Enter to finish):")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			token := strings.TrimSpace(scanner.Text())
			if token != "" {
				count++
			}
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
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.TrimSpace(line) != "" {
				count++
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}
	fmt.Println("All tokens must be in email:pass:token format.")
	fmt.Printf("Loaded %d token(s)", count)
	fmt.Println()
	fmt.Println()
	fmt.Println("\t[1] Convert to token format")
	fmt.Println("\t[2] Convert to token:pass format")
	fmt.Println("\t[3] Convert to token:email:pass format")
	fmt.Println()
	fmt.Println("What do you choose?")
	fmt.Print("-> ")
	var choice string
	fmt.Scan(&choice)

	fileOutput, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return
	}
	defer fileOutput.Close()

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := strings.TrimSpace(scanner.Text())
		tokenData := strings.Split(token, ":")
		email := tokenData[0]
		password := tokenData[1]
		token = tokenData[2]
		if choice == "1" {
			_, err = fileOutput.WriteString(token + "\n")
		} else if choice == "2" {
			_, err = fileOutput.WriteString(token + ":" + password + "\n")
		} else if choice == "3" {
			_, err = fileOutput.WriteString(token + ":" + email + ":" + password + "\n")
		} else {
			fmt.Println("Invalid choice.")
		}
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}
}
