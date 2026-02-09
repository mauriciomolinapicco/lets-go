// buffer channels: store multiple values at the same time without blocking

package main

import "fmt"


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


