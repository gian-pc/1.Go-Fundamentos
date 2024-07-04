package main

import "fmt"

func main(){

	result := sum(2)(3)
	fmt.Println(result) // 5

}

// función que retorna otra función
// sum: recibe un parámetro a de tipo int
// la función sum retorna una función que recibe un entero y retorna un entero
func sum(a int) func(int) int{
	return func(b int) int{
		return a + b
	}
}