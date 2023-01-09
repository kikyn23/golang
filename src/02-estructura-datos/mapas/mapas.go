package main

import (
	"fmt"
)

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

	//Nuevo Mapa
	estudiantes := make(map[string][]int)

	estudiantes["Alejandro"] = []int{13, 15, 16}
	estudiantes["Enrique"] = []int{12, 10, 19}

	fmt.Println(estudiantes)

	//Acceder a un elemento específico de la lista
	fmt.Println(estudiantes["Alejandro"][0]) // Imprime el valor 13, porque es la posición 0 del arreglo
}
