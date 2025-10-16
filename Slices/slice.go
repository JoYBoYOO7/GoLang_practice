package main

import "fmt"

func main() {
	//slice-> dynamic size
	//most used data structure in go
	//uninitialized slice is nil
	var numbers []int
	var nums = make([]int, 5)
	fmt.Println("Slice with make:", nums)
	fmt.Println("Capacity of slice:", cap(nums))

	// Assign values to the slice elements
	fmt.Println("Initial slice:", numbers)
	numbersSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", numbersSlice)

	// Append to a slice
	numbersSlice = append(numbersSlice, 6)
	fmt.Println("After appending 6:", numbersSlice)

	// Slicing a slice
	subSlice := numbersSlice[1:4]
	fmt.Println("Sub-slice from index 1 to 3:", subSlice)

	// Length and capacity of a slice
	fmt.Println("Length of slice:", len(numbersSlice))
	fmt.Println("Capacity of slice:", cap(numbersSlice))

	// Iterate over a slice using range
	for index, value := range numbersSlice {
		fmt.Println("Element at index", index, "is", value)
	}
	// Copying a slice
	copiedSlice := make([]int, len(numbersSlice))
	copy(copiedSlice, numbersSlice)
	fmt.Println("Copied slice:", copiedSlice)

	// Multidimensional slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Multidimensional slice:", matrix)

	// Modifying an element in the slice
	numbersSlice[0] = 10
	fmt.Println("After modifying first element:", numbersSlice)

}
