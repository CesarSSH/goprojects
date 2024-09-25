package arrays

import (
	"fmt"
)

func Arrays() {
	//I defined an array with only 5 data
	// var listOfSports = [5]string{"Basket", "Soccer", "Boxing", "Voley", "Golf"}

	// fmt.Println(listOfSports) //I printed the whole array
	// fmt.Println(listOfSports[3]) //I printed the third element of the array

	//Another way to create an array
	// languages := [2]string{"Php", "GO"}
	// fmt.Println(languages)
	// languages[0] = "Python"
	// fmt.Println(languages)

	//A slice doesn't have size, that's the difference with arrays
	planguages := []string{"Php", "GO", "Typescript"}
	fmt.Println(planguages)
	planguages[0] = "Python"
	fmt.Println(planguages)
	//I added a new programming language
	planguages = append(planguages, "Java")
	fmt.Println(planguages)

	//Ranges to get specific info from a range
	planguagestwo := planguages[1:3] //It takes the values 1, 2 only
	fmt.Println(planguagestwo)

	//Another way to get info from a range

	planguagesthree := planguages[2:] //It ignores the values 1, 2 and brings 3, 4....
	fmt.Println(planguagesthree)

	//Another way to get info from a range

	planguagesfour := planguages[:2] //It ignores the values 3, 4 and brings 1, 2
	fmt.Println(planguagesfour)
}
