package main

import "fmt"

func main(){
	// Por defecto Go asigna un string vacío "" a la variable a, cuado no se le asigna un valor

	var a string

	fmt.Printf("Tipo: %T, Valor: %q\n", a, a) // string, ""
}