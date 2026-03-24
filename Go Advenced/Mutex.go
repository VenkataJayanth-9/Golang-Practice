package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	noofgoroutines := 3
	wg.Add(noofgoroutines)

	increment := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for i := 0; i < noofgoroutines; i++ {
		go increment()
	}

	wg.Wait()
	fmt.Println("The value of the counter:", counter)
}

// type counter struct {
// 	mu    sync.Mutex
// 	count int
// }

// func (c *counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count++
// }

// func (c *counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.count
// }

// func main() {
// 	var wg sync.WaitGroup
// 	c := &counter{}

// 	numGoroutines := 10

// 	for i := 0; i < numGoroutines; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for j := 0; j < 1000; j++ {
// 				c.increment()
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Final Count:", c.getValue())
// }