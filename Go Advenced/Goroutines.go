package main

import (
	"fmt"
	"time"
)

func main(){

	var err error

	go func ()  {
		err = doWork()
	}()

	if err != nil {
		fmt.Println("Error : ", err)
	}else{
		fmt.Println("Successfully completed execution.")
	}

	fmt.Println("This is go beginning")
	go Hello()
	fmt.Println("Before 2 seconds")
	time.Sleep(2*time.Second)
	fmt.Println("After 2 seconds")

	go printLetters()
	go printNumbers()

	time.Sleep(20*time.Second)
	fmt.Println("This is the end hold your breath count to 10.")
}

func Hello() {
	fmt.Println("Hello")
	time.Sleep(2*time.Second)
}

func printNumbers() {
	for i:= 1; i<9; i++ {
		fmt.Println(i)
		time.Sleep(1*time.Second)
	}
}

func printLetters() {
	for _, i := range "abcde"{
		time.Sleep(3*time.Second)
		fmt.Println(string(i))
	}
}

func doWork () error {
	time.Sleep((1*time.Second))
	return fmt.Errorf("an error occured")
}