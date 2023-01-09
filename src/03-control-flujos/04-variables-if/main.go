package main

import "fmt"

func main() {

	if nombre, edad := "Enrique", 36; nombre == "Alejandro" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Println("Nombre", nombre, "Edad", edad)
	}

	//Obtener valor de mapa

	users := make(map[string]string)

	users["Enrique"] = "enrique.camelo@mmc.com"
	users["Camelo"] = "kikyn23@mmc.com"

	if correo, ok := users["Enrique"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}

}
