package main

import "fmt"

func main(){

	// reutilizando las funciones
	suma10 := sum(10)
	fmt.Println(suma10(2))
	fmt.Println(suma10(4))
	fmt.Println(suma10(1))
	fmt.Println(suma10(5))

}

// función que retorna otra función
// sum: recibe un parámetro a de tipo int
// la función sum retorna una función que recibe un entero y retorna un entero
func sum(a int) func(int) int{
	return func(b int) int{
		return a + b
	}
}