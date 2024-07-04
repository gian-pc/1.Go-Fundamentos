package main

import "fmt"

// pila
defer fmt.Println(1)
defer fmt.Println(2)
defer fmt.Println(3)

func main(){
	// Defer: La función defer nos permite diferir algo
	// diferir significa aplazar, que queremos aplazar la ejecución de una función
	// y hasta que momento queremos aplazar esa ejecución
	// hasta que la función en donde fue llamada el defer retorne

	// cuando se tiene multiples funciones diferidas se va agregar a la pila
	// una vez la función main retorne va ejecutar la funciones en la pila 1, 2 , 3
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	

}

// output: 1 2 3
