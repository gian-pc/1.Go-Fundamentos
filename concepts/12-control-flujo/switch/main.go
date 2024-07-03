package main

import "fmt"

func main(){
	character := "🦸"

	switch character{
	case "🦸":
		fmt.Println("es un superheroe")
	case "🦹":
		fmt.Println("es un supervillano")
	default:
		fmt.Println("es un personaje normal")
	}
}