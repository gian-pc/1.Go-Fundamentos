package main

import "fmt"

func main() {
	// operador de asignación de variable corta
	apple, banana, orange  := "manzana", "Banana", "Naranja"
	apple, lemon := "Manzana", "Lemon" // esto si es posible debido a que hay una variable nueva(lemon)

	fmt.Println(apple, banana, orange, lemon)

}
