package main

import (
	"fmt"
	"gopl.io/ch2/tempconv"
)

// set for tests
const absoluteZeroK = 0
const boilingF = 212

func main() {
	fmt.Println("Absolute zero:")
	fmt.Printf("Celsius: %v\n", tempconv.KToC(absoluteZeroK))
	fmt.Printf("Fahrenheit: %v\n", tempconv.KToF(absoluteZeroK))
	fmt.Printf("Kelvin: %v\n", absoluteZeroK)

	fmt.Println("Freezing:")
	fmt.Printf("Celsius: %v\n", tempconv.FreezingC)
	fmt.Printf("Fahrenheit: %v\n", tempconv.CToF(tempconv.FreezingC))
	fmt.Printf("Kelvin: %v\n", tempconv.CToK(tempconv.FreezingC))

	fmt.Println("Boiling:")
	fmt.Printf("Celsius: %v\n", tempconv.FToC(boilingF))
	fmt.Printf("Fahrenheit: %v\n", boilingF)
	fmt.Printf("Kelvin: %v\n", tempconv.FToK(boilingF))
}