package main

import "fmt"

func main(){
	/* TODO: Si se puede inferir el tipo si eliminamos el tipo de dato
		en este caso infiere "string" e "int"
	*/
	const os, domain= "Linux", 1
	fmt.Println(os, domain)
}