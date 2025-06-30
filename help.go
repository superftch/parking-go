package main

import "fmt"

// showHelp displays all available commands
func showHelp() {
	fmt.Println("AVAILABLE COMMANDS:")
	fmt.Println("==========================================")
	fmt.Println("create_parking_lot <capacity>")
	fmt.Println("park <car_number>")
	fmt.Println("leave <car_number> <hours>")
	fmt.Println("status")
	fmt.Println("help")
	fmt.Println("clear")
	fmt.Println("exit/quit")
	fmt.Println("")
	fmt.Println("PRICING = $10 for first 2 hours, $10 for each additional hour")
	fmt.Println("==========================================")
	fmt.Println()
}
