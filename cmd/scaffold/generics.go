package main

import "fmt"

// Constraint: tipos que admiten comparación
type Number interface {
	int | int64 | float64
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(3, 7))       // int
	fmt.Println(Max(10.5, 2.3))  // float64
	fmt.Println(Max[int64](9, 4))
}


