package main

import (
	"fmt"
	"time"
)

func main() {
	var expression int = 1
	var condition int = 1
	//basic switch
	switch expression {
	case condition:
		fmt.Println("Matched case 1")
	case 2:
		fmt.Println("Matched case 2")
	case 3:
		fmt.Println("Matched case 3")
	default:
		fmt.Println("Matched default")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	//expressionless switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	//type switch
	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		fmt.Printf("i is an int: %d\n", v)
	case string:
		fmt.Printf("i is a string: %s\n", v)
	default:
		fmt.Printf("i is of a different type: %T\n", v)
	}
}
