/*
***Generics in GO***

Generics allow you to write type-safe, reusable code that works with multiple data types without duplicating logic.
In simple terms: write once, use with many types, safely.

func Area[T int | float](l, b T) T{
	return l*b;
}

| Feature     | `interface{}` | Generics     |
| ----------- | ------------- | ------------ |
| Type safety | ❌ No          | ✅ Yes        |
| Errors      | Runtime       | Compile-time |
| Performance | Slower        | Faster       |
| Readability | Poor          | Better       |


*/

package main

import "fmt"

func Area[T int | float64](l, b T) T {
	return l*b
}

func main(){
	a := 5
	b := 6
	result := Area(a, b)
	c := 5.5
	d := 8.2
	res := Area(c, d)
	fmt.Printf("%.2f\n", res)
	fmt.Println(result)
}