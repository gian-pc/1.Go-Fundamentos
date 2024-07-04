package main

import (
	"fmt"
	"strings"
)

func main(){
	lower, upper := convert("GianPC")
	fmt.Println(lower, upper)
}

// funciones con multiples retornos - nombrados : se recomienda cuando las funciones son pequeñas
func convert(text string) (lower string, upper string){
	lower = strings.ToLower(text)
	upper = strings.ToUpper(text)

	return
}