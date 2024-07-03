package main

import "fmt"

func main(){
	
	character := "🦸"
	
	if  character == "🦸"{
		fmt.Println("es un superheroe")

	}else if character == "🦹"{
		fmt.Println("es un supervillano")

	}else{
		fmt.Println("es un personaje normal")

	}

	fmt.Println(character)
}