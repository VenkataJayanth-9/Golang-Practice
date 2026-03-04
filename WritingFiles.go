package main

import (
	"fmt"
	"os"
)

func main(){

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file.", file)
	}
	defer file.Close()
	//write data to file
	//Write is a method which takes only slice as arguments.
	data := []byte("Hello World!!!\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error!!!")
	}
	fmt.Println("Data is succesfully written")

	file1, err := os.Create("output1.txt")
	if err != nil {
		fmt.Println("Error creating the file")
		return
	}
	defer file.Close()
	_, err = file1.WriteString("Hello this is write string")

	if err != nil {
		fmt.Println("Error wrting to file.")
		return
	}

	fmt.Println("Written")
}