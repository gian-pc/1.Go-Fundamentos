package main

import "fmt"

func main(){

	division(100,10) // 10
	division(200,25) // 8
	division(34, 0) // panic
	division(124, 8) // no se ejecuta porque el programa entro en pánico y se detiene la ejecución

}

func division(dividend, divisor int){
	validateZero(divisor)
	fmt.Println(dividend / divisor)
}

func validateZero(divisor  int){
	if divisor == 0 {
		panic("😱 no puedes dividir por cero")
	}
}