package main

import (
	"fmt"
)

func main() {
	var numero int
	numero = 25
	fmt.Println(numero)
	numero = 40
	fmt.Println(numero)
	nombre := "Pedro"
	nombre = "Mauricio"
	nombre, numero = "Peter", 35
	fmt.Println(nombre, numero)
}
