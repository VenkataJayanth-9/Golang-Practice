package main

import (
	"fmt"
	"time"
)

func Workers(id int, jobs <- chan int, result chan <- int){
	for job := range jobs{
		fmt.Printf("Worker %d is doing work %d\n", id, job)
		time.Sleep(500*time.Millisecond)
		result <- 2*job
	}
}

func main(){
	jobs := make(chan int, 5)
	result := make(chan int, 5)

	for i := 1; i<=3; i++ {
		go Workers(i, jobs, result)
	}

	for z := 1; z <=5; z++{
		jobs<-z
	}

	close(jobs)

	for x := 0; x<5; x++{
		fmt.Println(<-result)
	}
}