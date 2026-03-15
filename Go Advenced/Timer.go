package main

import (
	"fmt"
	"time"
)

//========Timer
// func main(){
// 	fmt.Println("Starting the app...")
// 	timer := time.NewTimer(2*time.Second)

// 	fmt.Println("Before the channel")
// 	//blocking in nature
// 	stopped := timer.Stop()
// 	if stopped {
// 		fmt.Println("Timer stopped: ", stopped)
// 	}
// 	timer.Reset(time.Second)
// 	fmt.Println("Timer Reset")
// 	go func ()  {
// 		<-timer.C
// 	}()

// 	fmt.Println("Ending of the program")
// }

//========Time.After
// func longRunningOperation(){
// 	for i := range 20{
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main(){
// 	timeout := time.After(5 * time.Second)
// 	done := make(chan bool)

// 	go func ()  {
// 		longRunningOperation()
// 		done <- true
// 	}()

// 	select{
// 	case <-timeout:
// 		fmt.Println("Operation times out")
// 	case <- done:
// 		fmt.Println("Operation completed")
// 	}
// }


//==============Scheduling delayed operations

// func main(){
// 	timer := time.NewTimer(2 * time.Second)

// 	go func ()  {
// 		<- timer.C
// 		fmt.Println("Delayed Operation")
// 	}()

// 	fmt.Println("Waiting...")
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("End of the program")
// }

//================Handling multiple timers

func main(){
	timer1 := time.NewTimer(1*time.Second)
	timer2 := time.NewTimer(2*time.Second)

	select {
	case <- timer1.C:
		fmt.Println("Timer1 expired.")
	case <- timer2.C:
		fmt.Println("Timer2 expired.")
	}

	time.Sleep(5*time.Second)
}