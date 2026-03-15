// package main

// import "fmt"

// func producer (ch chan <- int) {
// 	for i := range 5{
// 		ch <- i
// 	}
// 	close(ch)
// }

// func filter (in <- chan int, out chan <- int ){
// 	for val := range in {
// 		if val % 2== 0{
// 			out <- val
// 		}
// 	}
// 	close(out)
// }

// func main(){
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go producer(ch1)
// 	go filter(ch1, ch2)

// 	for i := range ch2 {
// 		fmt.Println("Recived: ", i)
// 	}

// }

//============== Code - 2 use if cancel in context

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <- ctx.Done():
// 		return "Operation canceled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		}else{
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}
// }

// func main(){
// 	ctx := context.TODO()

// 	result := checkEvenOdd(ctx, 5)
// 	fmt.Println("Result with context.TODO(): ", result)

// 	ctx = context.Background()
// 	ctx, cancle := context.WithTimeout(ctx, 1 * time.Second)
// 	defer cancle()

// 	result = checkEvenOdd(ctx, 10)
// 	fmt.Println("Result from timeout context: ", result)

// 	time.Sleep(2*time.Second)
// 	result = checkEvenOdd(ctx, 15)
// 	fmt.Println("Last line of the code")
// }


//=========Storing the values in the conext and using them

package main

import(
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context){
	for{
		select{
		case <- ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working")
		}
		time.Sleep(500*time.Millisecond)
	}
}

func main(){
	ctx := context.Background()
	// ctx, cancle := context.WithTimeout(rootCtx, 2*time.Second)

	ctx, cancle := context.WithCancel(ctx)

	go func ()  {
		time.Sleep(2 * time.Second) //simulating a heavy task. consider this a heavy time consuming operation
		cancle()
	}()
	defer cancle()

	ctx = context.WithValue(ctx, "requestID", "12205857")
	go doWork(ctx)
	time.Sleep(2*time.Second)

	requestID := ctx.Value("requestID")

	if requestID != nil {
		fmt.Println("Request Id: ", requestID)
	}else {
		fmt.Println("No request ID found.")
	}
}