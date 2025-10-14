package main

import "fmt"

const PI = 3.14

func main() {
	const name string = "Vansh"
	const age int = 20
	const isStudent bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Value of Pi:", PI)

	//multiple constants
	const (
		country  = "India"
		language = "Go"
		version  = 1.18
	)
	fmt.Println("Country:", country)
	fmt.Println("Language:", language)
	fmt.Println("Version:", version)

	//constant expressions
	const radius = 5
	const area = PI * radius * radius
	fmt.Println("Area of circle with radius", radius, "is", area)

	//iota example
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
}
