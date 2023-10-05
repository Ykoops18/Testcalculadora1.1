package main

import (
	"fmt"
)

func main() {
	var valor1, valor2 float64
	var operador string

	fmt.Println("__________________________________________________")
	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&valor1)

	fmt.Print("Ingrese el tipo de operación que desea realizar (+, -, *, /): ")
	fmt.Scan(&operador)

	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&valor2)
}
