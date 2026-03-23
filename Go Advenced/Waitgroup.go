package main

import (
	"fmt"
	"sync"
	"time"
)

//Construction example

type Worker struct {
	ID int
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup){
	
}




//=============Example with channels
// func worker(id int, tasks <- chan int, result chan<-int, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Printf("WorkerID %d starting\n", id)
// 	time.Sleep(time.Second) //simulate some work
// 	result <- id+1
// 	fmt.Printf("WorkerID %d finished\n", id)
// }

// func main(){
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 5
// 	results := make(chan int, numJobs)
// 	tasks := make(chan int, numJobs)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, tasks, results, &wg)
// 	}

// 	for i := range numJobs{
// 		tasks<-i+1
// 	}

// 	close(tasks)

// 	//This makes non blocking 
// 	go func ()  {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Println("Result: ", result)
// 	}
// }

// ============ Basic example without using channels
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// wg.Add(1) This is the wrong practice where you should not write add method inside the function add should be written in only main function.
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers)

// 	//Launch workers
// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Last part of the code\n")
// }
