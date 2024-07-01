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
	// sin embargo un map nos va a devolver 2 valores un contenido y un boleano que nos va a decir si existe o  no una llave dentro del mapa
	content, ok := music["fake"]
	fmt.Println(content, ok) // vacío, false

}