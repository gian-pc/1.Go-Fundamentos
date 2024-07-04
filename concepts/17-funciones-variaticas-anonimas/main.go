package main

func main(){
	sum(2, 3)
	// ¿Que pasa si necesito llamar la función sum com más valores
	// esto no sería posible y se generaría un error
	sum(2,3,12)
	sum(2,3,12,1,24)

	
}

func sum(a, b int) int{
	return a + b
}