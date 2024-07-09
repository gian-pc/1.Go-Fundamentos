package main

import (
	"fmt"
)

func main(){
	// constraint comparable --> que viene por defecto en Go
	// me va permitir trabajar con restricciones de tipo
	// en donde solo se utiliza operadores de comparación (== !=)
	// y para ello no tendría que definir un tipo explicito como un entero o un float
	// porque simplemente la lógica que voy a tener dentro de mi función es para comparar algo

	fmt.Println( Includes([]string{"a","b","c"},"c")) // true
	fmt.Println( Includes([]string{"a","b","c"},"d")) // false
	fmt.Println( Includes([]int{1,12,24},24)) // true
}

// a esta función le pasamos un slice con multiples valores y le preguntamos que si el valor que queremos buscaar esta dentro del slice
func Includes[T comparable](list []T, value T) bool {
	for _, item := range list{
		if item ==  value{
			return true
		}
	}

	return false
}