package main

import "fmt"

func main() {

	dias := make(map[int]string)

	// fmt.Println(dias)

	//Agregar datos
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miércoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sábado"

	fmt.Println(dias)

	// Eliminar Elementos
	delete(dias, 1)
	fmt.Println(dias)
}
