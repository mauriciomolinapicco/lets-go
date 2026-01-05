package main

import "fmt"

func main() {
	var intArr [3] int32 // [0,0,0]
	fmt.Println(intArr[0])

	// memory location

	// it assings contiguous memory locations (beacuse int32 is 4 bytes each position of the array is 4 bytes apart)
	fmt.Println(&intArr[0])	
	fmt.Println(&intArr[1])	
	fmt.Println(&intArr[2])	

	// quick definition of array
	otroArr := [...]int32{1, 2, 3}
	fmt.Println(otroArr)

	// arrays are fixed size. CANNOT APPEND ELEMENTS
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))


	//another way of creating slices
	anotherSlice := make([]int32, 3, 5) // length 3, capacity 5
	// conviene definir la capacidad si mas o menos se que tan grande va a ser el slice
	//EVITA REALLOCATIONS Y COPIES
	anotherSlice[0] = 1
	anotherSlice[1] = 2
	anotherSlice[2] = 3
	fmt.Println(anotherSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(anotherSlice), cap(anotherSlice))

}
