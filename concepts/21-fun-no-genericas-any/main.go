package main

import "fmt"

type MyInt int

func main(){
	var num1 MyInt = 2
	var num2 MyInt = 6

	fmt.Println(sum(2,4,67))
	fmt.Println(sum[float32](2.0,4.6,67.1))
	fmt.Println(sum(num1, num2))
}

// Otra solución: constrains a partir de interface
type Number interface{
	~int | ~float64 | ~float32 | ~uint
}

func sum[T Number](nums ...T) T { 
	var total T
	for _, num := range nums{
		total += num
	}

	return total
}

