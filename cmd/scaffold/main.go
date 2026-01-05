package main

import "fmt"

func main() {
	var shornum int16 = 5
	var longnum int64 = 106376373

	var myUint uint = 34 // can store more numbers than int but only positive

	var float float32 = 34.567

	fmt.Println("Lets Go!!!!")

	var myString string = "Hello go" // one liner
	fmt.Println(myString)

	var anotherString string
	anotherString = ` this is a 
	multi line string`

	fmt.Println(len(anotherString))
	/* IMPORTANTE: len me devuelve la cantidad de BYTES, no de caracteres.
	Un caracter no ASCII puede ocupar mas de un byte
	*/
	fmt.Println(shornum, longnum, myUint, float)

	var myBool bool = true
	fmt.Println(myBool)

	/* 
	defaults: 0, false
	*/

	// var str1 = "Hello" // inferred type
	// str2 := "World" //shorter declaration with inferred type

	/* CONST */
	const myConst string = "I am a constant" //tengo que declararla en el momento con el valor. Constantes son inmutables

	printMe()
}



