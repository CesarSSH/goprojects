package datatype

import "fmt"

func DataT() {

	//Strings
	var myname string = "Sebastian" //First way
	var lastname = "Ruiz"           //Second way, I didn't declare the type of data. In this case, it's string.
	//Third way
	nickname := "Cezz"

	//Int
	var age int8 = 24 //-127 till 127 
	var year int16 = 2024
	var height float32 = 1.80
	var weight float32 = 70.0

	//Boolean
	var isSingle bool = false
	var isDeveloper bool = true

	fmt.Println("My name is ", myname)
	fmt.Println("My lastname is ", lastname)
	fmt.Println("My nickname is ", nickname)
	fmt.Println("What is the current year? ", year)
	fmt.Println("My age is ", age)
	fmt.Println("my height is ", height)
	fmt.Println("my weigth is ", weight)
	fmt.Println("Am I developer? ", isDeveloper)
	fmt.Println("Am I single? ", isSingle)

}
