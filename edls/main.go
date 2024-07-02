package main

import (
	"flag"
	"fmt"
)

func main(){
	// utilizando el paquete flag: nos va permitir hacer esos programas de lineas de comando
	// es un package de la librería standar

	// filter pattern
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "all files including hide files")
	flagNumberRecords := flag.Int("n", 0, "number of records") // flagNumberRecords

	// order flags
	hasOrderByTime := flag.Bool("t", false, "sort by time, oldest first")
	hasOrderBySize := flag.Bool("s", false, "sort by file size, smallest first")
	hasOrderReverse := flag.Bool("r", false, "reverese order while sorting")
	
	flag.Parse() // Parse va a mapear cada uno de los flags y los va a almacenar en flagPattern
	
	fmt.Println("pattern:",*flagPattern)
	fmt.Println("all:",*flagAll)
	fmt.Println("flagNumberRecords:",*flagNumberRecords)
	fmt.Println("hasOrderByTime:",*hasOrderByTime)
	fmt.Println("hasOrderBySize:",*hasOrderBySize)
	fmt.Println("hasOrderReverse:",*hasOrderReverse)
}

// para ejecutar ubicarse dentro de edls
// $ go mod init main: Para que cree el archivo "go.mod"
// para ejecutarlo usar: $ go run .
// filtrar solo png: $ go run . -p .png
// flag de ayuda: $ go run . -h

// utilizando flag -a: go run . -p .png -a // .png true
// utilizando flag -a: go run . -p .png // .png false

// records de 10: $ go run . -p .png -a -n 10

// flag -t: $ go run . -p .png -a -n 10 -t

// flag -s: $ go run . -p .png -a -n 10 -t -s

// flag -r: $ go run . -p .png -a -n 10 -t -s -r