package main

import (
	"fmt"
	"strings"
)

func main(){
	lower, upper := convert("GianPC")
	fmt.Println(lower, upper)
}

// funciones con multiples retornos - nombrados
func convert(text string) (lower string, upper string){
	lower = strings.ToLower(text)
	upper = strings.ToUpper(text)

	return
}