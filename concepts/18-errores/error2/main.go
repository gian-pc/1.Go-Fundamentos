package main

import (
	"errors"
	"fmt"
)

// creando nuestro primer error
// podemos crear un error con la función New del package errors
var errNotFound = errors.New("not found")
func main(){

	found, err := search("1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(found)

}

func search(key string) (string, error){
	return "", errNotFound // forzosamente estamos devolviendo el error
}