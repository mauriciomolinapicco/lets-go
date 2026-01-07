// POINTERS are a type 
package main
import "fmt"

func main() {
	var p *int32 = new(int32) // el new le devuelve una direccion de memoria
	var i int32

	fmt.Println("Valor de p:", p)   // direccion de memoria
	fmt.Println("Valor de *p:", *p) // dereferencia al valor (0) (preguntando lo que efectivamente hay en la direccion de memoria que tiene el puntero)
	fmt.Println("Valor de i:", i)    // valor por defecto (0)
}

