package main

import "fmt"

// make new parking lot
func NewParkingLot(capacity int) *ParkingLot {
	slots := make([]ParkingSlot, capacity)
	for i := 0; i < capacity; i++ {
		slots[i] = ParkingSlot{
			SlotNumber: i + 1,
			IsOccupied: false,
		}
	}
	return &ParkingLot{
		Capacity: capacity,
		Slots:    slots,
	}
}

// handle park
func (pl *ParkingLot) Park(registrationNumber string) {
	for i := 0; i < pl.Capacity; i++ {
		if !pl.Slots[i].IsOccupied {
			pl.Slots[i].IsOccupied = true
			pl.Slots[i].RegistrationNumber = registrationNumber
			fmt.Printf("Allocated slot number: %d\n", pl.Slots[i].SlotNumber)
			return
		}
	}
	fmt.Println("Sorry, parking lot is full")
}

// handle leave
func (pl *ParkingLot) Leave(registrationNumber string, hours int) {
	for i := 0; i < pl.Capacity; i++ {
		if pl.Slots[i].IsOccupied && pl.Slots[i].RegistrationNumber == registrationNumber {
			pl.Slots[i].IsOccupied = false
			slotNumber := pl.Slots[i].SlotNumber
			pl.Slots[i].RegistrationNumber = ""

			// calculate
			charge := calculateCharge(hours)

			fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n",
				registrationNumber, slotNumber, charge)
			return
		}
	}
	fmt.Printf("Registration number %s not found\n", registrationNumber)
}

// show status
func (pl *ParkingLot) Status() {
	fmt.Println("Slot No.\tRegistration No.")
	hasOccupied := false
	for i := 0; i < pl.Capacity; i++ {
		if pl.Slots[i].IsOccupied {
			fmt.Printf("%d\t\t%s\n", pl.Slots[i].SlotNumber, pl.Slots[i].RegistrationNumber)
			hasOccupied = true
		}
	}
	if !hasOccupied {
		fmt.Println("(No cars parked)")
	}
}

func calculateCharge(hours int) int {
	charge := 10 //base charge for 2 hour
	if hours > 2 {
		charge += (hours - 2) * 10
	}
	return charge
}
