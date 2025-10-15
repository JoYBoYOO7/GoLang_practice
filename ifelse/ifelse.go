package main

func main() {
	age := 18
	if age < 18 {
		println("Minor")
	} else if age == 18 {
		println("Adult")
	} else {
		println("Senior")
	}
	// go doesn't have ternary operator, use if-else instead

	var role = "admin"
	var hasPermission = true
	if role == "admin" && hasPermission {
		println("Access granted")
	} else {
		println("Access denied")
	}
}
