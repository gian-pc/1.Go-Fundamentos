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

	// ahora vamos a poblar nuestra estructura file creada en edls 
	// y lo vamos a poblar en un slice de archivos(file)
	fs := []file{}

	// leyendo los archivos del directorio que nos esta enviando el usuario
	for _, dir := range dirs{
		getFile(dir, false) // creamos la función getFile y le pasamos el directorio(dir) y un valor adicional que es para  indicar si ese archivo es oculto o no, temporalmente le vamos a decir que ningun archivo es oculto ya que va a variar según el SO
	}

}

// Implementando getFile
func getFile(dir fs.DirEntry, isHidden bool)(file, error){
	return file{}, nil // retornamos un archivo o un error

}
