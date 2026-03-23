package main

import "fmt"

func main(){
	str := "Challenge"
	mp := make(map[rune]int)

	for _, val := range str {
		mp[val]++
	}

	for i, val := range mp{
		fmt.Println(string(i))
		fmt.Println(val)
	}
}