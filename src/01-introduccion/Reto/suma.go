package main

import "fmt"

func suma(a, b int) int {
	return (a + b)
}

func main() {

	numero1 := 0
	numero2 := 0

	fmt.Print("Ingrese un valor el primer número: ")
	fmt.Scanln(&numero1)

	if numero1 < 0 {
		println("El número debe ser positivo")

	}

	fmt.Print("Ingrese un valor el segundo número: ")
	fmt.Scanln(&numero2)

	if numero2 < 0 {
		println("El número debe ser positivo")

	}

	resultado := suma(numero1, numero2)

	println("La suma es:", resultado)

}
