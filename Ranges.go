/*
"range is a keyword used with for loops to iterate over elements of a collection."
Supported types:
Arrays
Slices
Strings
Maps
Channels

-----Syntax of range loop
for index, value := range collection{
	//body
}

*/


package main
import "fmt"

func main(){

	//range with Arrays
	arr := [5]int{9, 18, 27, 36, 45}
	for index, value := range arr{
		fmt.Println(index, value)
	}
	//In above value is a copy of the elements in the array

	//ranges with slices
	s := []int{9, 18, 27}
	for index, value := range s{
		fmt.Println(index, value)
	}

	//range with strings
	str := "Hello Go"
	for index, value := range str {
		fmt.Println(index, value)
	}
	//It prints the ASCII values of the each and every charavter in the string.

	//ranges wih maps
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for index, value := range m {
		fmt.Println(index, value)
	}
}