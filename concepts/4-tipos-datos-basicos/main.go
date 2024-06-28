package main

import "fmt"

func main(){
	// No se pueden hacer operaciones con diferentes tipos de datos: uint8 - uint16

	var a uint8 = 200
	var b uint16 = 2500

	// operador _blank: sirve para poder nombrar una lógica y no nos de problema al ejecutarla
	_ = uint16(a) + b

	fmt.Printf("Tipo: %T, Valor: %v\n", a, a) // uint8, 200
}