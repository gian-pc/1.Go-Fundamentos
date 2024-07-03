package main

import "fmt"

func main(){
	animals := [5]string{"gorila", "perro", "gato", "ave", "elefante"}
	pets := animals[1:3] // perro, gato

	// agregamos un nuevo elemento loro y pasa algo extraño nuevamente
	// el arreglo inicial animal ya no subre ninguna modificación
	// la capacidad se vuelve 8
	pets = append(pets, "conejo", "tortuga", "loro")

	// Array[4]{"perro", "gato", "ave", "elefante"}
	// cuando alcanzamos la capacidad maxima ya no podemos seguir modificando el array y Go crea un nuevo array de 8
	// new Array[8]{"perro", "gato", "conejo", "tortuga", "loro"} // 5

	fmt.Println("animals: " , animals) // animals:  [gorila perro gato ave elefante]
	fmt.Println("pets: " , pets) // pets:  [perro gato conejo tortuga loro]
	fmt.Println("tamaño pets: ", len(pets)) // 5
	fmt.Println("capacidad pets: ", cap(pets)) // 8

}