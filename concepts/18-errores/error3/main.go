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

	found, err := search("3") // pudimos controlar el error correctamente --> No encuentra el valor en el mapa
	// comparando un error específico
	if err == errNotFound{
		fmt.Println("pudimos controlar el error correctamente")
		return
	}

	if err != nil {
		fmt.Println("search()", err)
		return
	}

	fmt.Println(found)

}

// search: recibe una llave de cualquiero cosa y la convertimos en un entero
// y luego vamos a ir a buscar en un map(food)
func search(key string) (string, error){
	num, err := strconv.Atoi(key) // recibimos key y lo convertimos a num
	if err != nil{
		return "", fmt.Errorf("trconv.Atoi(): %v", err)
	}

	emoji, err := findFood(num)
	if err != nil {
		return "", fmt.Errorf("findFood(): %v", err)
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

// output: search() findFood(): not found --> de esta manera puedo rastrear que función me produjo el error y así poder corregirlo rápidamente