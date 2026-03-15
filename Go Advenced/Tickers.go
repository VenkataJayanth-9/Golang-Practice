package main

import (
	"fmt"
	"time"
)


// ===========Basic Ticker 
// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()
// 	i := 1
// 	for range 5{
// 		i *= 2
// 		fmt.Println(ticker)
// 	}
// }


//============== periodic task

// func periodicTask() {
// 	fmt.Println("Tick Tick")
// }

// func main(){
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select{
// 		case<-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

//=============== Ticker stopping gracefully.

func main(){
	ticker := time.NewTicker(time.Second)
	stop := time.After(5*time.Second)

	defer ticker.Stop()
	
	for {
		select{
		case tick := <- ticker.C:
			fmt.Println("Tick at: ", tick)
		case <-stop :
			fmt.Println("Stopping ticker.")
			return
		}
	}
}