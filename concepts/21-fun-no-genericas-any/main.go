package main

import "fmt"

func main(){
	fmt.Println(sum(2, 5, 7))
	// si quisieramos sumar tipos de datos flotante
	fmt.Println(sum(2.4, 5.5, 7.2)) // no se podría y se generaría un error ya que nuestra función esta limitada a int
}

// posible solución:  genericos - any --> se generaría un error debido al alcance de any
// ya que estamos haciendo una operación invalida y que el operador + no esta definida sobre la variable total
// ya que total es de tipo any
// al decirlo que es de tipo any, le estas diciendo que puede recibir cualquier tipo de dato(bool, string, int)
// y eso no es lo que queremos, lo que queremos es sumar enteros y flotantes
// por lo tanto any no sería la solución
func sum(nums ...any) any {
	var total any
	for _, num := range nums{
		total += num
	}

	return total
}

