// POINTERS are a type 
package main
import "fmt"

func main() {
	var p *int32 = new(int32) // el new le devuelve una direccion de memoria
	var i int32

	fmt.Println("Valor de p:", p)   // direccion de memoria ej 0x1400011c004
	fmt.Println("Valor de *p:", *p) // dereferencia al valor (0) (preguntando lo que efectivamente hay en la direccion de memoria que tiene el puntero)
	fmt.Println("Valor de i:", i)    // valor por defecto (0)

	*p = 10
	fmt.Println("Valor de *p despues de asignar 10:", *p) // dereferencia al valor (10)

	p = &i //ahora el puntero apunta a la direccion de i
	*p = 25  //cambio el valor de i a traves del puntero
	fmt.Println("Valor de p despues de asignar la direccion de i:", i) // valor de i (25)


	//useful when we use large parameters so we dont have to copy the intire data

}

