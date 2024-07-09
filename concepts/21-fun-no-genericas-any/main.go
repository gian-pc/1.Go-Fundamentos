package main

import "fmt"

func main(){
	fmt.Println(sum(2, 5, 7))
	// si quisieramos sumar tipos de datos flotante
	fmt.Println(sum(2.4, 5.5, 7.2)) // no se podría y se generaría un error ya que nuestra función esta limitada a int
}

// solución:  genericos - any
func sum(nums ...int)int{
	var total int
	for _, num := range nums{
		total += num
	}

	return total
}

