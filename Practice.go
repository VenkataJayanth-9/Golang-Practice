package main

import(
	"sync"
	"fmt"
)

type counter struct{
	mu 		sync.Mutex
	count 	int
}

func (c *counter)increment(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter)getcounter()int{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main(){
	var wg sync.WaitGroup
	c := &counter{}
	noofgoroutines := 10
	
	for i:= 0; i<noofgoroutines; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			c.increment()
		}()
	}
	wg.Wait()
	fmt.Println("This is the value of the counter: ", c.getcounter())
}









