package main

import "fmt"

func main(){
	ch := make(chan int)
	produceOnly(ch)
	recieveOnly(ch)
}

func produceOnly(ch chan <- int){
		go func () {
		for i := range 5 {
			ch<-i
		}
		close(ch)
	}()
}

func recieveOnly(ch <- chan int){
	for value := range ch {
		fmt.Println("Value of the channel: ", value)
	}
}