package main

import (
	"fmt"
	"strings"
)

// Si queremos modificar un valor debemos trabajar con funciones que pasan parámetros por referencia
// para ello debemos usar punteros
func main(){
	name := "gian"
	toUpperCase(&name) // de puntero a string se hace con el operador de dirección &

	fmt.Println(name) // GIAN : Si se modificó
}

func toUpperCase(text *string){ // ya no queremos recibir un tipo de dato string sino puntero de esta manera obtengo la referencia de la variable name y lo que obtengo es una dirección de memoria
	*text = strings.ToUpper(*text) // debo trabajar con el operador de referenciación * para acceder al valor
}