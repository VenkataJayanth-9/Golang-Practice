package main

import "fmt"

func Palindrom(front int, back int, str string) bool {
	for front < back {
		if str[front] != str[back] {
			return false
		}
		front++
		back--
	}
	return true
}

func main(){
	str := "aappleneenelppaa"

	var front int;
	var back int;
	front = 0;
	back = len(str)-1

	result := Palindrom(front, back, str)

	fmt.Println(result)
}