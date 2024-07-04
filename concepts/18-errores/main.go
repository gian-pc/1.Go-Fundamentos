package main

import (
	"fmt"
	"strconv"
)

func main(){
	// los erroes: son una pieza fundamental de Go
	// ya que se decidió al momento de su diseño trabajar con errores y no con excepciones
	// ya que las excepciones en la mayoría de casos no son controladas por el desarrollador
	// la idea es que los errores siempre sean controlados al momento que se presentan
	// este es el motivo por el cual las funciones retornan multiples valores
	// usualmente veremos que las funciones retornan como último argumento un error
	// para que podamos validar si la función que ejecutamos produjo un error y así poderlo controlar

	num, err := strconv.Atoi("gian-pc") // la función Atoi del package strconv convierte de string a int
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(num) // strconv.Atoi: parsing "gian-pc": invalid syntax
}