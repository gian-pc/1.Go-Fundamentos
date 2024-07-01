package main

import "fmt"

func main(){
	animals := [5]string{"gorila", "perro", "gato", "ave", "elefante"}
	pets := animals[1:3] // perro, gato

	// agregamos un nuevo elemento tortuga y llegamos a la capacidad máxima
	pets = append(pets, "conejo", "tortuga")

	fmt.Println("animals: " , animals) // animals:  [gorila perro gato conejo tortuga]
	fmt.Println("pets: " , pets) // pets:  [perro gato conejo tortuga]
	fmt.Println("tamaño pets: ", len(pets)) // 4
	fmt.Println("capacidad pets: ", cap(pets)) // 4
}