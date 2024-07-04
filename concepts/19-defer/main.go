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

	// de esta  manera se puede implementar defer en nuestro codigo
	defer file.Close() // con esto me aseguro que siempre se cierre el archivo

	_, err = file.Write([]byte("Hello gophers")) // una vez creado el archivo voy a escribir contenido dentro del archivo test.txt
	if err != nil {
		fmt.Println(err)
		return
	}
}

// ejecutamos go run main.go y comprobamos si se a creado el archivo test.txt
// leemos el archivo creado: cat test.txt
