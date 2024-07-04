package main

import "fmt"

func main(){

	// funciones anónimas: son simplemente funciones que no tienen nombre

	greet := func()  {
		fmt.Println("👋Hola")
	}

	greet()
}