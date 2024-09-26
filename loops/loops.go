package loops

import (
	"fmt"
	"strings"
)

func Loops() {
	//GO uses only FOR for everything. I mean, it doesn't have WHILE, DO WHILE and stuff like that.
	//But we can use FOR to achieve the same functionality.
	// var total int = 0
	// for i := 1; i <= 10; i++ {
	// 	// fmt.Println("Number:",i)
	// 	// total += i
	// 	//Only plus the odd numbers
	// 	if i%2 != 0 {
	// 		total = total + i
	// 		fmt.Println("The total updated:", total)
	// 	}

	// }

	// myMap := map[string]string{ //String is the key and string is the value
	// 	"Colombia":  "Bogota",
	// 	"Venezuela": "Caracas",
	// 	"Brazil":    "Brasilia",
	// 	"Chile":     "Santiago",
	// }
	// //It's something like FOR EACH
	// for key, value := range myMap {
	// 	fmt.Println("The capital of ", key, "is :", value)
	// }

	//Implementing DO WHILE
	// var fruit string = "watermelon"

	// for{
	// 	if fruit == "orange" {
	// 		fmt.Println("The user chose orange...")
	// 		break

	// 	}
	// 	fmt.Println("Write your favorite fruit:")
	// 	fmt.Scan(&fruit)
	// 	fruit = strings.ToLower(fruit)
	// }

	//Implementing WHILE
	// var fruit string = "watermelon"
	var fruit string = ""

	for {

		fmt.Println("Type the right fruit:")
		fmt.Scan(&fruit)
		fruit = strings.ToLower(fruit)
		if fruit == "orange" {
			fmt.Println("The user chose orange...")
			break

		} else {
			fmt.Println("Incorrect data")
		}
	}

}
