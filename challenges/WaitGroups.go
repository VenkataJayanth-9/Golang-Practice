package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(i int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Working %d\n", i)
	time.Sleep(time.Second)
}

func main(){

	Jobs := 10
	var wg sync.WaitGroup

	wg.Add(Jobs)

	for i := 1; i<=10; i++{
		go Worker(i, &wg)
	}
	wg.Wait()
}