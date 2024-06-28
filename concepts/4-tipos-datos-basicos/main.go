package main

import "fmt"

func main(){
	// Por defecto Go asigna un 0 a la variable a, cuado no se le asigna un valor

	var a uint8

	fmt.Printf("Tipo: %T, Valor: %v\n", a, a) // uint8, 0
}