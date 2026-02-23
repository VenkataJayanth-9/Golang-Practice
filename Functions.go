/*
A function is a reusable block of code that performs a specific task.

Functions are first-class citizens

Can be assigned to variables

Can be passed as arguments

Can be returned from other functions

-----******--------
syntax of function 

func functionName(parameter1 data_type1, parameter2 data_type2) returnType{
	function_body
}

*/

package main

import "fmt"

//Simple function where we are adding two numbers.
//In this function we are taking 2 int parameters and we are retuning int
func add(a int, b int) int {
	return a+b
}

//No parameter and no return type function 
func greet() {
	fmt.Println("Good afternool!!")
}

//***very very imp***
// Multiple return values in a single function
func division(a int, b int) (int, error){
	if b == 0 {
		return 0, fmt.Errorf("invalid divsion")
	}
	return a/b, nil
}

//***very very imp*** 
//veriadic function
//A variadic function is a function that can accept a variable number of arguments of the same type.
// Declared using ... (ellipsis)
// Internally treated as a slice
// Allows flexible function calls
func sum(a ...int)int{
	total := 0
	for _,v := range a{
		total += v
	}
	return total
}

/*
***very very imp***
Anonymous Function
An anonymous function is a function without a name.
Defined inline
Can be assigned to variables
Can be passed as arguments
Can form closures
*/

/*
*****very very very imp*****
closures
A closure is a function value that captures and remembers variables from its surrounding lexical scope,
even after that scope has finished executing.
In Go, closures are created using anonymous functions.
*/

func counter() func() int {
	count := 0
	return func() int{
		count++
		return count
	}
}

func main(){
	result :=add(1, 2)
	fmt.Println(result)
	greet()
	fmt.Println("If we divide a and b")

	//In division function we are having two returns so we have to declear with 2 variables
	div_res, error := division(9, 3)
	if error != nil {
		fmt.Println(div_res)
		return
	}
	fmt.Println(div_res)

	//veriadic function call
	s1 :=sum(1, 2, 4)
	s2 :=sum(9, 18, 27, 36, 45)
	s3 := sum()

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	count := counter();
	fmt.Print(count)

} 