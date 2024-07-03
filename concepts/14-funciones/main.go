package main

import (
	"fmt"
	"strings"
)

// Las funciones por defecto van a trabajar con parámetros por valor
func main(){
	name := "gian"
	toUpperCase(name)

	fmt.Println(name) // gian : no se llega a modificar a mayúscula el nombre gian
}

func toUpperCase(text string){
	text = strings.ToUpper(text)
}