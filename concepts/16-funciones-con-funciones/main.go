package main

import "fmt"

func main(){
	nums := []int{2,12,23,98,21,79}

	// llamamos nuestra función filter
	resultado := filter(nums, menoresAVeinte)
	fmt.Println(resultado) // [98 79]

}

func mayoresACincuenta(i int) bool {
	return i > 50
}

// podemos generar otras funciones y reutilizar la función filter
func menoresAVeinte(i int) bool {
	return i < 20
}

// la función filter recibe un slice(nums) de enteros([]int)
// y va a recibir una función que vamos a llamar callback
// esta función tendrá por parámetro un valor de tipo entero
// y retornará un booleano
// finalmente la función filter va a retornar un slice de enteros []int que sería el resultado de filtrar los valores
func filter(nums []int, callback func(int) bool) []int{
	// en result almacenaremos los valores filtrados
	result := make([]int, 0, len(nums)) //[]int: creamos un slice de enteros con la función make, 0: no sabemos el tamaño de elementos que queremos filtrar, len(nums): pero si podemos establecer una capacidad máxima de slices para que Go no tenga que calcularlo 

	// iterar por el slice que estamos recibiendo
	for _, num := range nums{
		//vamos a pasar a la función callback cada elemento que estamos iterando
		//agregamos una condición: que va depender de lo que el usuario quiera filtrar talvés nums > 10 no es de interés saberlo
		//si se cumple esa condición, porque la función callback devuelve un booleano
		if callback(num){
			result = append(result, num) //vamos agregar el valor de ese número que estamos iterando al resultado
		}

	}
	// finalmente retornamos la variable result
	return result

}