package main

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func ()  {
		time.Sleep(2*time.Second)
		ch1 <- 1
	}()

	select {
	case msg := <-ch1 :
		fmt.Println("ch1", msg)
	case msg := <-ch2 :
		fmt.Println("ch2", msg)
	default:
		fmt.Println("Default")
	}

	fmt.Println("End of the program")
}