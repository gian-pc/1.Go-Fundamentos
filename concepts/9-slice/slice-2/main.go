package main

import "fmt"

func main(){
	animals := [5]string{"gorila", "perro", "gato", "ave", "elefante"}
	pets := animals[1:3] // perro, gato

	//para agregar un elemento a pets se usa append, pero pasa algo extraño. Se modifica ave x conejo
	pets = append(pets, "conejo")

	fmt.Println("animals: " , animals) // animals:  [gorila perro gato conejo elefante]
	fmt.Println("pets: " , pets) // pets:  [perro gato conejo]
	fmt.Println("tamaño pets: ", len(pets)) // 3
	fmt.Println("capacidad pets: ", cap(pets)) // 4
}