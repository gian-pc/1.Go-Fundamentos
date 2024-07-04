package main

import "fmt"

func main(){
	num := 4
	// la función diferida hace que Go antes de agregar a la pila evalua los argumentos de mi función en este caso num = 4
	defer fmt.Println("imprime", num) // 1. se agrega esta función a la pila     5. se ejecuta la funcion diferida

	num = 10 // 2. actualizamos el número 10
	fmt.Println(num) // 3. se imprime el número 10
	
	// 4. la función main retorna

}

// output: 10, imprime 4
