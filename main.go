package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("PARKING LOT MANAGEMENT SYSTEM")
	fmt.Println("=================================")

	// check if input is file
	if len(os.Args) == 2 {
		filename := os.Args[1]
		runFromFile(filename)
		return
	}

	// cli mode
	fmt.Println("CLI Mode - Type commands or 'help' for assistance")
	fmt.Println("Type 'exit' to quit")
	fmt.Println("")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("parking-lot> ")

		if !scanner.Scan() {
			break
		}

		command := strings.TrimSpace(scanner.Text())
		if command == "" {
			continue
		}

		processCommand(command)
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
