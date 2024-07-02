package main

import "fmt"

// las estructuras son similares a las clases en otros lenguajes de programación
// usualmente la declaramos a nivel de package para que esten disponible para todas las funciones

type Person struct{
	Name string
	Age uint8
	HasChildren bool
}

func main(){

	// como creamos instancias de Persona
	// primera forma recomendable
	gian := Person{
		Name: "Gian",
		Age: 38,
		HasChildren: false,
	}

	// segunda forma
	jose := Person{"Jose", 33, true}

	// tercera forma
	andres := Person{Age:20} //si solo quiero agregar el dato Age

	fmt.Printf("%+v\n", gian)
	fmt.Printf("%+v\n", gian.Name) //si solo quiero acceder a Name
	fmt.Printf("%+v\n", jose)
	fmt.Printf("%+v\n", andres)
}