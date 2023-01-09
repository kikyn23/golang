package main

import "fmt"

func main() {
	// arrays
	var numeros [5]int

	numeros[0] = 1
	numeros[1] = 3
	numeros[2] = 4
	numeros[3] = 6
	numeros[4] = 9

	fmt.Println(numeros)

	//Arrays con valores
	nombres := [3]string{"Enrique", "Sandra", "Alejandro"}

	fmt.Println(nombres)

	//Arrays con tamaño definido por el contenido
	colores := [...]string{"Amarillo", "Azul", "Naranja"}

	fmt.Println(colores, len(colores))

	//Arrays con tamaño definido por el contenido
	monedas := [...]string{
		0: "Peso",
		2: "Dollar",
		4: "Euro"}

	fmt.Println(monedas, len(monedas))

	//Subarrays
	subMonedas := monedas[0:3]

	fmt.Println(subMonedas)
}
