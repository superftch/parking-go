package main

// ParkingSlot represents a single parking slot
type ParkingSlot struct {
	SlotNumber         int
	RegistrationNumber string
	IsOccupied         bool
}

// ParkingLot represents the entire parking lot
type ParkingLot struct {
	Capacity int
	Slots    []ParkingSlot
}

// Global parking lot variable
var parkingLot *ParkingLot
