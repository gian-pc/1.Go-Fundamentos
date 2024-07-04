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
// podriamos modificar esta función con los retornos nombrados
func sum(nums  ...int) (total int){

	for _, num := range nums {
		total += num
	}

	return total
}