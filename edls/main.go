package main

import (
	"flag"
	"fmt"
)

func main(){
	// utilizando el paquete flag: nos va permitir hacer esos programas de lineas de comando
	// es un package de la librería standar

	// -p: me va filtrar la salida
	flagPattern := flag.String("p", "", "filter by pattern")
	flag.Parse() // Parse va a mapear cada uno de los flags y los va a almacenar en flagPattern
	fmt.Println(*flagPattern)
}

// para ejecutar ubicarse dentro de edls
// $ go mod init main: Para que cree el archivo "go.mod"
// para ejecutarlo usar: $ go run .
// filtrar solo png: $ go run . -p .png
// flag de ayuda: $ go run . -h
