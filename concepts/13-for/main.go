package main

import "fmt"

func main(){
	// iterando el slice literal

	for i, v := range []string{"🍕", "🍔", "🍎", "🌭"}{
		fmt.Println("indice:", i, "valor", v)
	}
}