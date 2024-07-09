package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)


func main(){
	// si quiero utilizar otro  operador de comparación > < >= <= 
	

	fmt.Println( Includes([]string{"a","b","c"},"c")) // true
	fmt.Println( Includes([]string{"a","b","c"},"d")) // false
	fmt.Println( Includes([]int{1,12,24},24)) // true

	fmt.Println(filter([]int{2,12,23,98,21,79}, menoresAVeinte)) // [2 12]
}

func menoresAVeinte(i int) bool {
	return i < 20
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

func filter[T constraints.Ordered](nums []T, callback func(T) bool) []T{

	result := make([]T, 0, len(nums))

	for _, num := range nums{
		if callback(num){
			result = append(result, num)
		}
	}
	return result

}