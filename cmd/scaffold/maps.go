package main

import "fmt"

func main() {
	// literal
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}

	/* 
	IMPORTANTE. Si llamo a una key que no existe GO devuelve el VALOR POR DEFAULT del tipo de dato que sea
	*/

	//go tambien devuelve un segundo objeto opcional BOOLEANO que indica si la key existe o no
	var v, ok = m1["c"]
	if ok {
		fmt.Printf("Value: %d\n", v)
	}

	//recorrer un map
	for key, value := range m1 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	var mySlice []int = []int{1,2,3,4,5}
	//recorrer un slice
	for index, value := range mySlice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}



}
