package main

import (
	"fmt"
)

// scope
var razakira = "golden"

func main() {
	var (
		user      = "admin"
		mail, rol = "admin@server.com", "admin"
	)
	var numero1 = 2

	fmt.Println(user, mail, rol, numero1)

	// scope
	imprimir()
}

func imprimir() {
	fmt.Println("La raza de Kira es: " + razakira)
}
