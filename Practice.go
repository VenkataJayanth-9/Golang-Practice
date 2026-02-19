package main
import "fmt"

func main(){
	//Arrays.
	var array[5] int
	fmt.Println("Hello World")
	array = [5]int{1, 2, 4, 5, 6}
	fmt.Println(array)

	//Slices.
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)

	//Struct.
	type Person struct{
		Name string
		Age int
	}

	p := &Person{
		Name: "Jayanth",
		Age: 23,
	}

	fmt.Println(p.Name)
	fmt.Println(p.Age)

	//Nested Structs
	type Address struct{
		City string
		Pin int
	}

	type User struct{
		Name string
		Address Address
	}
}