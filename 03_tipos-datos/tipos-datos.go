package main

import (
	"fmt"
)

func main() {
	// Enteros con signo (-+)
	var entero8 int8
	var entero16 int16
	var entero32 int32
	var entero64 int64

	// Enteros sin signo (-+)
	var uentero8 uint8
	var uentero16 uint16
	var uentero32 uint32
	var uentero64 uint64

	//  Alias
	var enteroByte byte // alias para uint8
	var enteroRune rune // alias para int32

	fmt.Println("// Enteros con signo (-+)")
	fmt.Println(entero8)
	fmt.Println(entero16)
	fmt.Println(entero32)
	fmt.Println(entero64)

	fmt.Println("// Enteros sin signo (-+)")
	fmt.Println(uentero8)
	fmt.Println(uentero16)
	fmt.Println(uentero32)
	fmt.Println(uentero64)

	fmt.Println("// Alias")
	fmt.Println(enteroByte)
	fmt.Println(enteroRune)

}
