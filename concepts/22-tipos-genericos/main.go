package main

import "fmt"

type Product[T uint | string] struct{
	ID	 		 T
	Description	 string
	Price	 	 float64
}

func main(){

	//esto es util para generar estructuras de forma generica
	// caso de ejemplo: si tenemos una DB en postgres con id entero y luego queremos pasarlo
	// a mongoDB con un id string "6f9b13ce-d9e8" se tendría que generar otra estructura

	product1 := Product[uint]{ID: 1, Description: "Zapatos", Price: 30}
	product2 := Product[string]{ID: "6f9b13ce-d9e8", Description: "Zapatos", Price: 30}

	fmt.Println(product1)
	fmt.Println(product2)

}