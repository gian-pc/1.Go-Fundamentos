package main

import "fmt"

func main(){
	// Por defecto Go asigna un false a la variable a, cuado no se le asigna un valor
	// el valor 0 de bool siempre será cero

	var a bool

	fmt.Printf("Tipo: %T, Valor: %v\n", a, a) // bool, false
}