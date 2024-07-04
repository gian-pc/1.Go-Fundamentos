package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errNotFound = errors.New("not found")

var food = map[int]string{
	1: "🍕",
	2: "🍔",
}

func main(){

	found, err := search("1") // 🍕 --> encuentra la pizza en el mapa
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(found)

}

// search: recibe una llave de cualquiero cosa y la convertimos en un entero
// y luego vamos a ir a buscar en un map(food)
func search(key string) (string, error){
	num, err := strconv.Atoi(key) // recibimos key y lo convertimos a num
	if err != nil{
		return "", err // retorna vacío porque no vamos a poder encontrar la 🍕, 🍔 y el error que se esta produciendo al no poder convertir key a un número
	}

	emoji, err := findFood(num)
	if err != nil {
		return "", err
	}

	return emoji, nil

}

func findFood(id int) (string, error){
	value, ok := food[id] // buca el id en el mapa y devuelve el valor del mapa como 🍕 o 🍔 y ok si fue encontrado esa llave
	if !ok { // sino existe ese valor en el mapa devolvemos el error antes creado errNotFound
		return "", errNotFound // devolvemos "" porque no econtramos ningún valor
	}

	return value, nil // si lo encontramos retornamos el emoji y nil(por que no encontramos ningún error)
}