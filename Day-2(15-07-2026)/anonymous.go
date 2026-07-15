package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello Golang")
	}()
	// func(name string) {
	// 	fmt.Println("Hello", name)
	// }("chandini")

}