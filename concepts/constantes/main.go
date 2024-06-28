package main

import "fmt"

// es recomendable que las constantes esten a nivel de package
const os, domain= "Linux", "ed.team"

func main(){
	
	fmt.Println(os, domain)
}