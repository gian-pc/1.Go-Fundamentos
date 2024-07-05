package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	
	flag.Parse()

	path := flag.Arg(0) 

	dirs, err := os.ReadDir(path) // nos permite leer los archivos de una dirección específica(path)
	if err != nil{
		panic(err)
	}
	fmt.Println(dirs)

	
}
// terminal: $ go run . /home/gian/Desktop/testFiles
// output: [- README.md - gian.png - go1.22.4.linux-amd64.tar.gz d myFolder/]