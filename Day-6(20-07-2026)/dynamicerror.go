package main

import "fmt"

func main() {

	name,err := "student.txt"

	// err := fmt.Errorf("Cannot open file: %s", name)//Dynamic message
	// if err!=nil{
		return fmt.Errorf("opening config file failed: %w", err)
	// }

	// fmt.Println(name)

}