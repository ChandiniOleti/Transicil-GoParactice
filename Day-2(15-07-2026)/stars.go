package main

import "fmt"

func main() {

	num := 10

	p := &num

	fmt.Println("Value of num:", num)

	fmt.Println("Address of num:", &num)

	fmt.Println("Pointer p:", p)

	fmt.Println("Value using pointer:", *p)

}