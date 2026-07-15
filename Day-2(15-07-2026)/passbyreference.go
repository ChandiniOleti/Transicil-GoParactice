package main

import "fmt"

func change(x *int) {
	*x = 100
}

func main() {
	num := 10
	p := &num
	change(p)
	fmt.Println(num)
}