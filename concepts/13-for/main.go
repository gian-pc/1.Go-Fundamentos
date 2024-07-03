package main

import "fmt"

func main(){
	// modificando nuestro slice
	numbers := []uint8{2,4,6,8}

	for i := range numbers{
		numbers[i] *= 2
	}
	fmt.Println(numbers)
}