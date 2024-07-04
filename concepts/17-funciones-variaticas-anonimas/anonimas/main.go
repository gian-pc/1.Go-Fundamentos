package main

import "fmt"

func main(){

	// funciones anónimas: son simplemente funciones que no tienen nombre

	func(name string)  { // le pasamos parámetros
		fmt.Println("👋Hola", name)
	}("Gian") // le pasamos el valor a la variable name

	
}