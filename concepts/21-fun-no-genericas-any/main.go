package main

import "fmt"

func main(){
	PrintListString("a", "b", "c")
	// problema: que  pasa si quiero imprimir números
	PrintListInt(1, 2, 3)
}

func PrintListString(list ...string){
	for _, item := range list{
		fmt.Println(item)
	}
}

// posible solución: se repite codigo
func PrintListInt(list ...int){
	for _, item := range list{
		fmt.Println(item)
	}
}