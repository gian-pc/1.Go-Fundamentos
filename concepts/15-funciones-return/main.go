package main

import (
	"fmt"
	"strings"
)

func main(){
	lower, upper := convert("GianPC")
	fmt.Println(lower, upper)
}

// funciones con multiples retornos
func convert(text string) (string, string){
	lower := strings.ToLower(text)
	upper := strings.ToUpper(text)

	return lower, upper
}