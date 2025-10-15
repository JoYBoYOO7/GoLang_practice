package main

//for only construct in Go
import "fmt"

func main() {
	// for init; condition; post { }
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for condition { }
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// for { } - infinite loop
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println(k)
		k++
	}

	// for range - iterate over array, slice, map, string, channel
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	str := "hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}
}
