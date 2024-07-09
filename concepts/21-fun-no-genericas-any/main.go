package main

import "fmt"

type MyInt int

func main(){
	var num1 MyInt = 2
	var num2 MyInt = 6

	fmt.Println(sum[int](2,4,67))
	fmt.Println(sum[float64](2.2,4.6,67.1))
	fmt.Println(sum[MyInt](num1, num2))
}

// - tipos de datos definidos --> enteros, flotantes, strings, booleanos, etc.
// - se pueden también crear nuestros propios tipos de datos basados en eso datos existentes
func sum[T ~int | float64 ](nums ...T) T {
	var total T
	for _, num := range nums{
		total += num
	}

	return total
}

