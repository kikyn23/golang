package main

import "fmt"

func main() {
	//Slicen
	numeros := []int{1, 2, 3}

	fmt.Println(numeros)

	//Agregar datos
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)
	fmt.Println(numeros)
}
