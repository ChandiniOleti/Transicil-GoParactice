package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Open source file which is already exisits
	source, err := os.Open("student.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer source.Close()

	// Create destination file 
	destination, err := os.Create("backup.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destination.Close()

	// Copy data from student to backup files
	bytesCopied, err := io.Copy(destination, source)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Copy Successful")
	fmt.Println("Bytes Copied:", bytesCopied)
}