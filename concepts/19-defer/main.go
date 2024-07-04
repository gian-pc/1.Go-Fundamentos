package main

import "fmt"

func main(){
	// Defer: La función defer nos permite diferir algo
	// diferir significa aplazar, que queremos aplazar la ejecución de una función
	// y hasta que momento queremos aplazar esa ejecución
	// hasta que la función en donde fue llamada el defer retorne

	defer fmt.Println(3) // defer: me va aplazar la ejecución del número 3
	fmt.Println(1)
	fmt.Println(2)
	

}

// output: 1 2 vuelve a el return de la función main y ejecuta los valores diferidos(3)
