package main

import "fmt"

func valorIva(valor, impuesto float64) float64 {
	return valor * impuesto
}

func valorVenta(valorVenta, impuesto float64) float64 {
	return valorVenta + impuesto
}

func main() {

	var valor, iva, total float64

	fmt.Print("Ingrese un valor de la venta: ")
	fmt.Scanln(&valor)

	iva = valorIva(valor, 0.18)

	total = valorVenta(valor, iva)

	println("El valor de la venta es:", valor)
	println("El impuesto es:", iva)
	println("El valor total es:", total)

}
