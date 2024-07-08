package main

import "fmt"

func main(){
	PrintList("a", "b", "c")
	PrintList(1, 2, 3)
}

// solución:  genericos - any
func PrintList(list ...any){
	for _, item := range list{
		fmt.Println(item)
	}
}