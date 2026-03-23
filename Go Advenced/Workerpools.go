package main

import (
	"fmt"
	"time"
)

type ticketRequest struct{
	personID 	int
	numTickets 	int
	cost 		int
}

//simulate processing of ticket requests
func ticketProcessor(request <- chan ticketRequest, results chan <- int){
	for req := range request{
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost %d\n", req.numTickets, req.personID, req.cost)
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func main(){
	numRequest := 5
	prices := 5
	ticketRequests := make(chan ticketRequest, numRequest)
	ticketResults := make(chan int)

	// start ticket processor/worker
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	//send ticket requests
	for i := range numRequest {
		ticketRequests <- ticketRequest{personID: i+1, numTickets: (i+1)*2, cost: (i+1)*prices}
	}
	close(ticketRequests)

	for range numRequest{
		fmt.Printf("Ticket for personID %d processed successfully!\n", <-ticketResults)
	}
}

// ============Basics of workerpools
// func worker(id int, tasks <- chan int, results chan <- int){
// 	for task := range tasks {
// 		fmt.Printf("Worker %d is processing task %d\n", id, task)
// 		//Simulate work
// 		time.Sleep(time.Second)
// 		results <- task*2
// 	}
// }

// func main(){
// 	numWorkers := 3 // no of goroutines
// 	numJobs := 10

// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	//Create workers
// 	for i := range numWorkers {
// 		go worker(i, tasks, results)
// 	}

// 	//Send values to the tasks channels
// 	for i := range numJobs{
// 		tasks <- i
// 	}
// 	close(tasks)

// 	for range numJobs{
// 		result := <- results
// 		fmt.Println("Results: ", result)
// 	}
// }