/* 
enable go routines that pass informacion
- hold data
- thread safe
- listen for data
*/

package main

import "fmt"

func fakeMain() {
	var c = make(chan int) //chan keyword for channels
	c <- 1  //its like it has an implicit array
	var i = <- c //the value 1 gets popped out of the channel c and assigned to i
	fmt.Println(i)

	//it will never run because there is no goroutine sending data to the channel
}

func main() {
	var c = make(chan int)
	go process(c)
	for i:= range c {
		fmt.Printf("Recibiendo valor %v \n", i)
	}
}

func process(c chan int) {
	defer close(c) //close the channel when function ends
	//defer significa: hace esto al final de la funcion
	for i:=0; i<5; i++ {
		fmt.Printf("Enviando valor %v\n", i)
		c <- i //send value to channel
	}
	fmt.Println("Process ended")
}

//SIEMPRE que se envia un dato c <- i se bloquea hasta que alguien reciba <-c


