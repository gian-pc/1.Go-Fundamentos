package main

import (
	"flag"
	"fmt"
	"io/fs"
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
		f, err := getFile(dir, false) // la función getFile retorna un archivo o u error
		// controlando el error
		if err != nil{
			panic(err)
		}
		
		fs = append(fs, f) // agregamos el archivo al slice
	}
	fmt.Println(fs) // imprimimos el contenido del slice fs

}

// Implementando getFile
func getFile(dir fs.DirEntry, isHidden bool)(file, error){
	// ahora vamos a implementar esta función file, lo único que va hacer es poblar la estructura file es decir todos esos campos que previamente incluimos
	// para ello lo primero que vamos hacer es tener información del archivo que esta en dir
	info, err := dir.Info() // Info nos va proveer información acerca del archivo para poder poblar nuestra estructura
	if err != nil{
		return file{}, fmt.Errorf("dir.Info(): %v", err) // en el caso que se presente un error devuelve un archivo vacío y el error
	}

	// llenamos la estructura
	f := file{
		name:				 dir.Name(),
		isDir:				 dir.IsDir(),
		isHidden:			 isHidden,
		userName:			 "gian-pc",
		groupName:			 "dev",
		size:				 info.Size(),
		modificationTime:	 info.ModTime(),
		mode:				 info.Mode().String(), // el mode nos prove info de los permisos de lectura, escritura, etc. de los archivos 
	}



	return f, nil // retornamos un archivo o un error

}
