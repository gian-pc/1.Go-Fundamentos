package main

import "fmt"

func main(){
	fmt.Println(sum[int](2,4,67))

}

// entoces es aquí donde entran los parámetros de tipo
// estos parámetros de tipo nos van a permitir definir restricciones para trabajar con multiples tipos de datos ya sean funciones o en tipos como estructura
// entonces cual sería la restricción en la función suma, pues recibir solo datos de tipo enteros y flotantes
// los parámetros de tipo se van a definir en corchetes
// existen 3 tipos de restricciones o constrains
	// - constrains arbitrario
	// - constrains de unión de elementos
	// - constrains de aproximación de elementos

func sum[T int](nums ...T) T {
	var total T
	for _, num := range nums{
		total += num
	}

	return total
}

