package main

import "fmt"

// Pero la mejor opción es el agrupamiento de constantes
const (
	os = "Linux"
	domain = "ed.team"
)

// las constantes a comparación de las variables no necesitan ser utilizadas para su compilación
const (
	Jan = iota + 1
	Feb
	Mar
	Abr
	May
	Jun

)

func main(){
	
	fmt.Println(Jan,Feb, Mar, Abr, May, Jun) // 1 2 3 4 5 6
}