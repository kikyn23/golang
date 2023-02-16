package main

import (
	"fmt"
	"paquetes/models"
)

func main() {
	// mensajes.Hola()
	// mensajes.Imprimir()

	// cua1 := figuras.Cuadrado{Lado: 8}
	// figuras.Medidas(&cua1)

	// cir1 := figuras.Circulo{Radio: 8}
	// figuras.Medidas(&cir1)

	p1 := models.Persona{}

	p1.Constructor("Enrique", 37)
	fmt.Println(p1.GetNombre())
	p1.SetNombre("Alejandro")
	fmt.Println(p1)

}
