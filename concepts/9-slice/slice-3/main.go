package main

import "fmt"

func main(){
	
	
	//pets := []string{"perro", "gato"}
	//pets := make([]string,0,3)
	//pets = append(pets, "conejo", "tortuga", "loro")

	var pets []string


	fmt.Println("pets: " , pets) //[]
	fmt.Println("tamaño pets: ", len(pets)) //0
	fmt.Println("capacidad pets: ", cap(pets)) //0
	fmt.Println("valor cero: ", pets == nil) //true

}