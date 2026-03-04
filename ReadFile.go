package main

import (
	"bufio"
	"fmt"
	"os"
	"text/scanner"
)

func main(){
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}
	defer func () {
		file.Close()
		fmt.Println("Closing the file.")
	}()
	fmt.Println("File opened successfully.")

	//Read the content of the file by opening the file.
	//Reading the file is in the form of byte slice.

	// data := make([]byte, 1024)
	// _, err = file.Read(data)

	// if err != nil {
	// 	fmt.Println("Error while reading the", err)
	// 	return
	// }

	// fmt.Println(string(data))

	scanner := bufio.NewScanner(file) // creates the instance of the scanner. Scanner reads from a file 
	//Read line by line

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading the file.")
		return
	}
}