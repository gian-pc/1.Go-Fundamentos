package main

import "fmt"

func main(){
	fmt.Println(sum[int](2,4,67))
	fmt.Println(sum[float64](2.2,4.6,67.1))
}

// - constrains de unión de elementos
func sum[T int | float64](nums ...T) T {
	var total T
	for _, num := range nums{
		total += num
	}

	return total
}

