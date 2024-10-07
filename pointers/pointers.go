package pointers

import "fmt"

func Pointers() {

	// Es una variable que nos permite acceder a la dirección en memoria de otra variable

	//El valor por defecto de un puntero es nulo
	value := 4
	pointer := &value

	fmt.Println(pointer)  //Returna un formato hexagecimal 0xc00008c098
	fmt.Println(*pointer) // Permite saber el valor que contiene en el formato hexagecimal osea 4 en este caso
	pointer = nil
}
