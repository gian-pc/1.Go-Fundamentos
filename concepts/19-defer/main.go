package main

import "fmt"

func main(){
	// Defer: La función defer nos permite diferir algo
	// diferir significa aplazar, que queremos aplazar la ejecución de una función
	// y hasta que momento queremos aplazar esa ejecución
	// hasta que la función en donde fue llamada el defer retorne

	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)

}

// output: 1 2 3