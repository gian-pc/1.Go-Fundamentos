package main

import "fmt"

func main(){
	// arrays
	var letters [3]string
	letters[0] = "a"
	letters[1] = "b"
	letters[2] = "c"

	// arrays literales
	num := [...]int{1,2,3,4,5}

	fmt.Println(letters)
	fmt.Println(num)
}