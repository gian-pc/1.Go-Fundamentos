package main

import "fmt"

func main(){
	
	// otra forma de crear slices
	pets := []string{"perro", "gato"}
	fmt.Println("pets: " , pets) 
	fmt.Println("tamaño pets: ", len(pets)) 
	fmt.Println("capacidad pets: ", cap(pets)) 

}