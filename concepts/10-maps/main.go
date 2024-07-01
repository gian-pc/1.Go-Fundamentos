package main

import (
	"fmt"
)

func main(){
	//make: funcion preconstruida
	//tipo de dato: map
	music := make(map[string]string)
	music["guitar"] = "guitarra"
	music["violin"] = "violin"

	fmt.Println(music)

	// maps de forma literal
	tech := map[string]string{
		"computador":"computador",
		"mouse":"mouse",
	}

	fmt.Println(tech)

	// eliminar un elemento de un map
	delete(tech, "mouse")
	fmt.Println(tech)

	// leer elementos de un map. Sino encuentra el valor devuelve un string vacío
	fmt.Println(music["guitar"])

}