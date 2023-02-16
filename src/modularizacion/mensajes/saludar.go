package mensajes

import "fmt"

/*
Funciones o variable que incien en mayúscula pueden ser accedidas desde otras partes del programa
Funciones o variables en minúscula, no pueden ser accesibles desde otras partes del código
*/

func Hola() {
	fmt.Println("Hola desde el paquete de mensajes")
}

const mensajes = "Mensaje guardado en una constante"

func Imprimir() {
	fmt.Println(mensajes)
}
