package main

import "fmt"

func main(){
	greet("Gian", "Paucar")
}

func greet(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}