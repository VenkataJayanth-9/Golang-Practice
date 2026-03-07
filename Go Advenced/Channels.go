package main

import "fmt"

func main(){
	greeting := make(chan string)
	greetingString := "Good Morning"

	go func ()  {
		greeting <- greetingString
	}()

	recevier := <- greeting
	fmt.Println(recevier)
}