package main

import (
	"fmt"
	"os"
)
func main() {

	file, err := os.OpenFile(
		"student.txt",
		os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString("\nNew Line Added")

	fmt.Println("Data Appended")
}