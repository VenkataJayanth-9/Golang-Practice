package main

import "fmt"

func main(){
	str := "golang"
	var newstr string
	for i:=len(str)-1; i>=0; i--{
		newstr += string(str[i])
	}
	fmt.Println(newstr)
}