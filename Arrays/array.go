package main

import "fmt"

func main() {
	// Create an array of integers with a length of 5
	var numbers [5]int
	var names [3]string

	// Assign values to the array elements
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"

	// Print the entire array
	fmt.Println("Array:", numbers)
	fmt.Println("Initial string array:", names)

	// Access and print individual elements
	fmt.Println("First element:", names[0])
	fmt.Println("Second element:", names[1])
	fmt.Println("Third element:", names[2])
	fmt.Println("First element:", numbers[0])
	fmt.Println("Second element:", numbers[1])

	// Get the length of the array
	fmt.Println("Length of array:", len(numbers))
	fmt.Println("Length of array:", len(names))

	// Iterate over the array using a for loop
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index", i, "is", numbers[i])
	}
	for i := 0; i < len(names); i++ {
		fmt.Println("Element at index", i, "is", names[i])
	}

	// Iterate over the array using a range loop
	for index, value := range numbers {
		fmt.Println("Element at index", index, "is", value)
	}
	for index, value := range names {
		fmt.Println("Element at index", index, "is", value)
	}
	// Declare and initialize an array in one line
	cities := [3]string{"New York", "Los Angeles", "Chicago"}
	fmt.Println("Cities:", cities)

	// Multidimensional array
	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	fmt.Println("Matrix:", matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Element at [%d][%d] is %d\n", i, j, matrix[i][j])
		}
	}
}
