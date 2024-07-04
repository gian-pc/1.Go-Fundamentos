package main

import "fmt"

func main(){
	
	// ¿Que pasa si necesito llamar la función sum com más valores
	fmt.Println(sum(2)) // 2
	fmt.Println(sum(2, 3)) // 5
	fmt.Println(sum(2,3,12)) // 17
	fmt.Println(sum(2,3,12,1,24)) // 42
}
// SOLUCIÓN: función variática que me va permitir recibir n argumetos
func sum(nums  ...int) int{
	var total int
	for _, num := range nums {
		total += num
	}

	return total
}