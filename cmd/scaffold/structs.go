package main

import "fmt"

// Basic struct definition
type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner   ownerInfo
}

type ownerInfo struct {
	name    string
	address string
}

// Electric engine struct
type electricEngine struct {
	kwhPer100Miles uint8
	kwhInBattery   uint8
	owner          ownerInfo
}

// Interface definition - defines a contract
type engine interface {
	milesLeft() uint8
	getOwner() ownerInfo
}

func main() {
	// Struct initialization - named fields
	var myGasEngine gasEngine = gasEngine{
		mpg:     30,
		gallons: 15,
		owner: ownerInfo{
			name:    "John Doe",
			address: "123 Main St",
		},
	}

	fmt.Println("=== Gas Engine ===")
	fmt.Println("MPG:", myGasEngine.mpg)
	fmt.Println("Gallons:", myGasEngine.gallons)
	fmt.Println("Owner Name:", myGasEngine.owner.name)
	fmt.Println("Owner Address:", myGasEngine.owner.address)
	fmt.Println("Miles left:", myGasEngine.milesLeft())

	// Struct initialization - positional (less readable)
	myElectricEngine := electricEngine{
		30, // kwhPer100Miles
		50, // kwhInBattery
		ownerInfo{
			"Jane Smith",
			"456 Oak Ave",
		},
	}

	fmt.Println("\n=== Electric Engine ===")
	fmt.Println("kWh per 100 miles:", myElectricEngine.kwhPer100Miles)
	fmt.Println("kWh in battery:", myElectricEngine.kwhInBattery)
	fmt.Println("Miles left:", myElectricEngine.milesLeft())

	// Using interfaces for polymorphism
	fmt.Println("\n=== Interface Polymorphism ===")
	var engines []engine = []engine{myGasEngine, myElectricEngine}

	for i, eng := range engines {
		fmt.Printf("Engine %d - Miles left: %d, Owner: %s\n", 
			i+1, eng.milesLeft(), eng.getOwner().name)
	}

	// Pointer receiver example
	fmt.Println("\n=== Pointer Receiver ===")
	myGasEngine.refuel(10)
	fmt.Println("After refueling, gallons:", myGasEngine.gallons)

	// Value receiver vs pointer receiver
	fmt.Println("\n=== Value vs Pointer Receiver ===")
	myGasEngine.setMPGValue(25) // Won't change original (value receiver)
	fmt.Println("After setMPGValue(25), mpg:", myGasEngine.mpg)

	myGasEngine.setMPGPointer(25) // Will change original (pointer receiver)
	fmt.Println("After setMPGPointer(25), mpg:", myGasEngine.mpg)
}

// Value receiver method - receives a copy
func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

// Value receiver method
func (e gasEngine) getOwner() ownerInfo {
	return e.owner
}

// Pointer receiver method - can modify the struct
func (e *gasEngine) refuel(gallons uint8) {
	e.gallons += gallons
}

// Value receiver - won't modify original
func (e gasEngine) setMPGValue(mpg uint8) {
	e.mpg = mpg // Only modifies the copy
}

// Pointer receiver - will modify original
func (e *gasEngine) setMPGPointer(mpg uint8) {
	e.mpg = mpg // Modifies the original struct
}

// Electric engine implements the engine interface
func (e electricEngine) milesLeft() uint8 {
	return (e.kwhInBattery * 100) / e.kwhPer100Miles
}

func (e electricEngine) getOwner() ownerInfo {
	return e.owner
}