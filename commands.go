package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// processCommand processes a single command
func processCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "create_parking_lot":
		handleCreateParkingLot(parts)
	case "park":
		handlePark(parts)
	case "leave":
		handleLeave(parts)
	case "status":
		handleStatus()
	case "help":
		showHelp()
	case "exit", "quit":
		handleExit()
	case "clear":
		handleClear()
	default:
		fmt.Printf("Unknown command: %s\n", parts[0])
		fmt.Println("Type 'help' to see available commands")
	}
}

// handleCreateParkingLot handles the create_parking_lot command
func handleCreateParkingLot(parts []string) {
	if len(parts) != 2 {
		fmt.Println("Usage: create_parking_lot <capacity>")
		fmt.Println("   Example: create_parking_lot 6")
		return
	}

	capacity, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Invalid capacity:", parts[1])
		return
	}

	if capacity <= 0 {
		fmt.Println("Capacity must be greater than 0")
		return
	}

	parkingLot = NewParkingLot(capacity)
	fmt.Printf("Created a parking lot with %d slots\n", capacity)
}

// handlePark handles the park command
func handlePark(parts []string) {
	if len(parts) != 2 {
		fmt.Println("Usage: park <car_number>")
		fmt.Println("Example: park KA-01-HH-1234")
		return
	}

	if parkingLot == nil {
		fmt.Println("Please create parking lot first using: create_parking_lot <capacity>")
		return
	}

	parkingLot.Park(parts[1])
}

// handleLeave handles the leave command
func handleLeave(parts []string) {
	if len(parts) != 3 {
		fmt.Println("Usage: leave <car_number> <hours>")
		fmt.Println("   Example: leave KA-01-HH-1234 4")
		return
	}

	if parkingLot == nil {
		fmt.Println("Please create parking lot first using: create_parking_lot <capacity>")
		return
	}

	hours, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Invalid hours:", parts[2])
		return
	}

	if hours <= 0 {
		fmt.Println("Hours must be greater than 0")
		return
	}

	parkingLot.Leave(parts[1], hours)
}

// handleStatus handles the status command
func handleStatus() {
	if parkingLot == nil {
		fmt.Println("Please create parking lot first using: create_parking_lot <capacity>")
		return
	}
	parkingLot.Status()
}

// handleExit handles the exit/quit command
func handleExit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

// handleClear handles the clear command
func handleClear() {
	// Clear screen command
	fmt.Print("\033[H\033[2J")
}
