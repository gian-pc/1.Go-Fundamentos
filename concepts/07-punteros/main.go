package main

import "fmt"

func main(){

	// Puntero: Variables que almacenan la dirección en memoria de un valor

	var color string = "red"
	var pointerColor *string // declarando un puntero, le estamos diciendo a Go que vamos almacenar una dirección en memoria y no un valor

	pointerColor = &color // almacernar la dirección en memoria de la variable color en pointerColor
	*pointerColor = "blue" // modificando desde nuestro puntero a la variable color


	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v\n", color, color, &color) // Tipo: string, Valor: red, Dirección: 0xc000014070
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", pointerColor, pointerColor, *pointerColor) // Tipo: *string, Valor: 0xc000014070, Desreferenciación: blue




}