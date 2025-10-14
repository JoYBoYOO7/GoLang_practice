package main

import "fmt"

func main() {
	//explicit declaration
	var a int = 1
	fmt.Println(a)
	var b string = "Hello"
	fmt.Println(b)
	var c float32 = 1.5
	fmt.Println(c)
	var d bool = true
	fmt.Println(d)
	//type inference
	var name = "Vansh"
	fmt.Println(name)
	//shorthand declaration
	tag := "JoyBoy"
	fmt.Println(tag)
}
