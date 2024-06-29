package main

import "fmt"

func main(){
	// slice: son apuntadores a un array, no poseen datos

	things := [6]string{"1","2", "3", "a", "b", "c"}
	num := things[:3] // no es necesario el valor inicial si es el 0
	letters := things[3:] //[a,b,c] no es necesario el valor final si es el último valor
	all := things[:]
	letters[1] = "B" // modificando el array [a,B,c], al modificar cambia también el array inicial things debido a que son direcciones de memoria


	fmt.Println(things)
	fmt.Println(num)
	fmt.Println(letters)
	fmt.Println(all)
}