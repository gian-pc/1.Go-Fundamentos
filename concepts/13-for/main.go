package main

import "fmt"

func main(){
	// for indefinido: se usa para trabajar con sockets y este a la escucha indefinidamente
	i := 1
	for {
		if i == 6{
			break
		}
		fmt.Println(i)
		i++
	}
}