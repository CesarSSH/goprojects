package conditions

import "fmt"

func Conditions() {

	// var age int8 = 70

	// if age >= 65 {
	// 	fmt.Print("I'm a senior")
	// } else if age >= 18 {
	// 	fmt.Print("I'm an adult")
	// } else {
	// 	fmt.Print("I'm a kid")
	// }

	// var age int8 = 56
	// var privilege bool = true

	// if age >= 18 && privilege {
	// 	fmt.Print("Access allowed, you can come in")
	// } else if age <= 17 && privilege {
	// 	fmt.Print("Access denied")
	// } else {
	// 	fmt.Print("Access allowed, you are a senior")
	// }

	//This is a very simple structure for SWITCH statement
	var fruit string

	fmt.Print("Write a string value: ")
	fmt.Scan(&fruit)

	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "lemon":
		fmt.Println("It's a lemon")
	default:
		fmt.Println("It's something else")

	}

}
