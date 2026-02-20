/*
Go maps 
A map is a built in reference data type that stores data as key-value pairs.

Keys are unique
Values are associated with keys
Lookup, insertion and deletion are efficient

---Syntax to create maps in go
map[keyType]valueType


--Valid key types
Keys must be comparable:
like int, string, bool, struct, pointers
but not slices, maps, functions

--value types
Any type is allowed

*/

package main
import "fmt"

func main(){
	fmt.Println("This file is about go maps");
	var m map[int]int
	/*The above syntax to create map allowed reading but causes runtime panic while writing.*/

	mm := make(map[int]int)
	/*Allocates memory 
	  Ready for use
	*/

	mmm := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m)
	fmt.Println(mm)
	fmt.Println(mmm)

	//How to access element of map
	value := mmm["a"]
	fmt.Println(value)

	//Comma-ok Idiom
	value, ok := mmm["a"]
	fmt.Println(ok)

	//Update and Insertion in map is same
	mmm["a"] = 9
	fmt.Println(mmm)

	//deleting element in map
	delete(mmm, "a")
	fmt.Println(mmm)

	//Iteration Over Maps
	//Iteration order is not guaranteed
	//Never rely on order
	for key, value := range mmm{
		fmt.Println(key, value)
	}

	//Find the length of the map
	a := len(mmm)
	fmt.Println(a)

}