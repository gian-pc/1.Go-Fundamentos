package main

import "fmt"

func main(){
	// operadores aritmeticos: () * / % + - en orden de jerarquía

	var a = (2 + 3)* 2
	fmt.Println(a) //10

	// operadores de asignación: = += -= *= /= %=
	var b int = 5
	b += 2
	fmt.Println(b)//7

	// Expresión vs Declaración
	// Expresión: una expresión es una parte de código que me produce un valor (línea 8)
	// Declaración: Es solo una instrucción del lenguaje que realiza  una acción
	// esta declaración puede incrementar (++) o decrementar (--) la variable en 1

	var c int = 7
	c++
	fmt.Println(c) //8

}