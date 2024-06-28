package main

import "fmt"

// Pero la mejor opción es el agrupamiento de constantes
const (
	os = "Linux"
	domain = "ed.team"
)

func main(){
	
	fmt.Println(os, domain)
}