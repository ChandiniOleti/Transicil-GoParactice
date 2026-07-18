package main

import (
	"fmt"
	"os"
)

func main() {

	err := os.WriteFile(
		"exp1.txt",
		[]byte("Hello Chandini"),
		0644,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File Written Successfully")
}