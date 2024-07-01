package main

import "fmt"

func main(){
	
	
	//pets := []string{"perro", "gato"}
	pets := make([]string,0,3)
	pets = append(pets, "conejo", "tortuga", "loro")

	fmt.Println("pets: " , pets) 
	fmt.Println("tamaño pets: ", len(pets)) 
	fmt.Println("capacidad pets: ", cap(pets)) 

}