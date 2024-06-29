package main

import "fmt"

func main(){
	// operadores comparación: > < == != >= <=
	fmt.Println(4 < 6)// true

	// operadores lógicos: && ||
	var age uint = 5
	fmt.Println("Es Adulto?: ", age >= 18 && age <= 60)

	// operador lógico Unario: !
	fmt.Println(!(4==4))


}