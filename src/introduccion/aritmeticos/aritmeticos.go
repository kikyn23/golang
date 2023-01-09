package main

import "fmt"

func main() {
	a := 20
	b := 10

	//Suma
	result := a + b

	//Resta
	result = a + b
	fmt.Println("Resta: ", result)

	//Multiplicacion
	result = a * b
	fmt.Println("Multiplicacion: ", result)

	//División
	result = a / b
	fmt.Println("División: ", result)

	//División con decimailes
	var division float64 = 3.0 / 2.0
	fmt.Println("División decimal: ", division)

	//Módulo
	result = 3 % 2
	fmt.Println("Modulo: ", result)
}
