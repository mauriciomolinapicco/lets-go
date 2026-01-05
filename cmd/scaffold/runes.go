package main

import (
	"fmt"
	"strings"
)

func main() {

	// A rune is an alias for int32 and represents a Unicode code point
	// reminder: len of a string counts bytes, not characters
	var runeArr = []rune("Hello, 世界")
	var indexed = runeArr[7]

	fmt.Printf("The rune at index 7 is: %v\n", indexed)

	for i, v := range runeArr {
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}

	//runes are unicode point numbers that represents characters
	// -----
	//runes are imutable. so when we concatenate it copies

	var strSlice = []string{"Hello", " ", "世界"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	fmt.Printf("Concatenated string: %s\n", strBuilder.String())
}