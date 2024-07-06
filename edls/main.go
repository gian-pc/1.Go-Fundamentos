package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	
	flag.Parse()

	path := flag.Arg(0) 

	// Que pasa sino enviamos el path - hay que validarlo antes que se envie un string vacío
	if path == ""{
		path = "." // sino se envía un path entonces por defecto se mostrará el path actual
	}

	dirs, err := os.ReadDir(path) // nos permite leer los archivos de una dirección específica(path)
	if err != nil{
		panic(err)
	}

	// leyendo los archivos del directorio que nos esta enviando el usuario
	for _, dir := range dirs{
		fmt.Println(dir.Name())
	}

	
}
// terminal: $ go run . 
// output:
	// edls.go
	// go.mod
	// main.go