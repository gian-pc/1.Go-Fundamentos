package main

import "fmt"

func main(){
	// iterando mapas
	food := map[string]string{
		"pizza": "🍕",
		"hamburguer": "🍔",
		"apple": "🍎",
		"hotdog": "🌭",
	}
	

	for key, value := range food{
		fmt.Println("key", key, "valor", value)
	}
	
}