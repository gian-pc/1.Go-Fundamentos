package main

import (
	"fmt"
	"os"
)

func main(){
	// en la práctica esto para que me funciona
	// esto es bastante util para limpiar recursos
	// cerrar archivos
	// cerrar conexiones de red
	// cerrar controladores de base de datos

	file, err :=  os.Create("test.txt") // nos va permitir generar un archivo en nuestro sistema operativo
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.Write([]byte("Hello gophers")) // una vez creado el archivo voy a escribir contenido dentro del archivo test.txt
	if err != nil {
		file.Close() // si se presenta un error y no podamos escribir se tendría que cerrar el archivo y así limpiar recursos
		fmt.Println(err)
		return
	}

	file.Close() // una vez escrito en el archivo test.txt lo voy a cerrar
}


// ejecutamos go run main.go y comprobamos si se a creado el archivo test.txt
// leemos el archivo creado: cat test.txt
