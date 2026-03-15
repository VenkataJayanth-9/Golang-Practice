package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("Hello Good Morning!!!")
	// ch := make(chan int)

	// =========Non Blocking code for receiving the value to the channel
	// go func ()  {
	// 	ch<-1
	// }()

	// time.Sleep(2*time.Second)

	// select{
	// case msg := <-ch:
	// 	fmt.Println("Message: ", msg)
	// default:
	// 	fmt.Println("No message to display.")
	// }

	//==========Non blocking sent function for a channel
	// go func ()  {
	// 	fmt.Println("This is case 1: ", <-ch)
	// }()

	// time.Sleep(2*time.Second)
	// select {
	// case ch <- 1:
	// 	fmt.Println("This is the case 1")
	// default:
	// 	fmt.Println("There are no recivers.")
	// }

	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case msg := <-data:
				fmt.Println("Message recieved: ", msg)
			case <-quit:
				fmt.Println("Quit the channel")
				return
			default:
				fmt.Println("Waiting...")
				time.Sleep(500*time.Millisecond)
			}
		}
		
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}
}
