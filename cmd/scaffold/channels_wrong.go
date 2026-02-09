package main

import "fmt"

func main() {
	var c = make(chan int)
	c <- 1

	// nunca va a pasar de aca porque se bloquea hasta que alguien lo lea
	// estan hechos para usarse con rutinas
	var i = <- c
	fmt.Println(i)
}
