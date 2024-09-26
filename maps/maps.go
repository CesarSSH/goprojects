package maps

import "fmt"

func Maps() {
	//This is the structure for a MAP
	myMap := map[string]string{ //String is the key and string is the value
		"Colombia":  "Bogota",
		"Venezuela": "Caracas",
		"Brazil":    "Brasilia",
		"Chile":     "Santiago",
	}

	fmt.Println("My map is:", myMap)
	//Go doesn't bring the same order for the list of values as I declared

	fmt.Println("What is the capital of Brazil?", myMap["Brazil"])

	//Add new item in my map
	myMap["Argentina"] = "Buenos Aires"
	fmt.Println("My updated map is:", myMap)

	//Remove an item
	delete(myMap, "Chile")
	fmt.Println("My updated map with a removed item is:", myMap)


	fmt.Println("What is the capital of Brazil?", myMap["Chile"]) //It returns empty


}
