package main

import (
	"fmt"
	"goprojects/operations"
)

func main() {
	fmt.Println("Hello, World!")
	// I call the Plus function and I save the result in a new variable
	result := operations.Plus(3, 5)

	// Print the result
	fmt.Println("Total: ", result)

}
