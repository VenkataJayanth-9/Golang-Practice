package main

import (
	"fmt"
	"time"
)

// func Ping(ping <- chan string, pong chan <- string){
// 	for {
// 		msg := <-ping
// 		fmt.Println(msg)
// 		time.Sleep(500*time.Millisecond)
// 		pong <- "pong"
// 	}
// }

// func Pong(ping chan <- string, pong <- chan string){
// 	for{
// 		msg := <-pong
// 		fmt.Println(msg)
// 		time.Sleep(500*time.Millisecond)
// 		ping <- "ping"
// 	}
// }

func main(){
	ping := make(chan string)
	pong := make(chan string)

	go func(){
		for{
			select{
	case msg := <-ping:
		fmt.Println(msg)
		time.Sleep(500*time.Millisecond)
		pong<-"pong"
	case msg := <-pong:
		fmt.Println(msg)
		time.Sleep(500*time.Millisecond)
		ping<-"ping"
	}
		}
	}()

	ping <- "ping"
	time.Sleep(5 * time.Second)
}