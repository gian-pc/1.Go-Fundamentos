package main

import "fmt"

func main(){
	// iterando el slice
	food := []string{"🍕", "🍔", "🍎", "🌭"}

	for i, v := range food{
		fmt.Println("indice:", i, "valor", v)
	}
}