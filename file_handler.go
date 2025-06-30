package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// runFromFile runs commands from a file
func runFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Printf("Reading commands from file: %s\n", filename)
	fmt.Println("================================")

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		if command != "" && !strings.HasPrefix(command, "#") { // Skip empty lines and comments
			fmt.Printf("[Line %d] Command: %s\n", lineNumber, command)
			processCommand(command)
			fmt.Println() // command split by enter
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Println("Finished processing file")
}
