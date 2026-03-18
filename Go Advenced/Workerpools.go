package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <- chan int, results chan <- int){
	for task := range tasks {
		fmt.Printf("Worker %d is processing task %d\n", id, task)
		//Simulate work
		time.Sleep(time.Second)
		results <- task*2
	}
}

func main(){
	numWorkers := 3 // no of goroutines
	numJobs := 10

	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//Create workers
	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	//Send values to the tasks channels
	for i := range numJobs{
		tasks <- i
	}
	close(tasks)

	for range numJobs{
		result := <- results
		fmt.Println("Results: ", result)
	}
}