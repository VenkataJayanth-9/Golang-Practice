package main

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <- chan int, result chan <- int){
	for job := range jobs {
		fmt.Printf("The worker %d is working on %d\n", id, job)
		time.Sleep(500*time.Millisecond)
		result <- 2*job
	}
}

func main(){
	files := []int{120, 450, 300, 200, 150, 600, 700}

	jobs := make(chan int, 7)
	result := make(chan int, 7)

	for i := 1; i<=3; i++{
		go Worker  (i, jobs, result)
	}

	for _, val := range files{
		jobs <- val 
	}
	close(jobs)

	for z := 0; z<5; z++ {
		fmt.Println(<-result)
	}
}