package main

import "fmt"

func main(){
	// No se pueden hacer operaciones con diferentes tipos de datos: uint8 - uint16

	var a uint8 = 200
	var b uint16 = 2500

	// solución para poder realizar la operación suma: 
	// casting 'a', pero nunca va a cambiar el tipo de dato de 'a' uint8
	c := uint16(a) + b

	fmt.Printf("Tipo: %T, Valor: %v\n", c, c) // uint16, 2700
	fmt.Printf("Tipo: %T, Valor: %v\n", a, a) // uint8, 200
}