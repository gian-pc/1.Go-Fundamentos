package main

import "fmt"

type MyInt int

func main(){
	var num1 MyInt = 2
	var num2 MyInt = 6

	fmt.Println(sum(2,4,67))
	fmt.Println(sum(2.2,4.6,67.1))
	fmt.Println(sum(num1, num2))
}

// por ende si se quiere que la función suma sea de forma generica soporte todos esos tipos de datos
// una posible solución sería incrementar esos tipos de datos en los parámetros
func sum[T ~int | ~float64 | ~float32 | ~uint](nums ...T) T { // en este caso solo estamos restringiendo int y float64 sin embargo hay mas tipos de datos numericos
	var total T
	for _, num := range nums{
		total += num
	}

	return total
}

