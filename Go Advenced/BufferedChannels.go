package main

import "fmt"

func main(){
	a := make(chan int, 2)
	a <- 1
	a <- 2
	fmt.Println("1st value of the buffer: ", <-a)
	fmt.Println("2nd value of the buffer: ", <-a)
	a <- 3
	fmt.Println(<-a)
}