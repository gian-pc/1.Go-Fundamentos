package main

import "fmt"

func main(){
	PrintList("a", "b", "c")
}

func PrintList(list ...string){
	for _, item := range list{
		fmt.Println(item)
	}
}