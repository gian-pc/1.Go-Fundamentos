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
}