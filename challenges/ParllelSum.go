package main

import (
	"fmt"
)


//============Using Wait Groups



//============Using channels

func Sum(array []int, ch chan <- int) {
	sum := 0
	for _, v := range array {
		sum += v
	} 
	ch<-sum
}

func main(){
	array := []int{1, 2, 3, 4, 5, 6}
	mid := len(array)/2
	ch := make(chan int)
	
	go Sum(array[:mid], ch)
	go Sum(array[mid:], ch)

	sum1 := <-ch
	sum2 := <-ch

	close(ch)

	result := sum1+sum2
	fmt.Println(result)
}



//=============Simple method
// func main(){
// 	//create an array
// 	arr := []int{1, 2, 3, 4, 5, 6}
// 	var sum1 int
// 	var sum2 int
// 	go func(){
// 		for i:=0; i<len(arr)/2; i++{
// 			sum1 = sum1+arr[i]
// 		}
// 	}()

// 	go func ()  {
// 		for i:=(len(arr)/2); i<len(arr); i++{
// 			sum2 = sum2+arr[i]
// 		}
// 	}()

// 	time.Sleep(2*time.Second)
// 	result := sum1+sum2

// 	fmt.Println(sum1)
// 	fmt.Println(sum2)

// 	fmt.Println(result)
// }