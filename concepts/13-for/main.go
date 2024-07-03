package main

import (
	"fmt"
)

func main(){
	// iterando sobre un string

	for i, v := range "Gian" {
		fmt.Println("indice", i, "value", string(v))
	}

}