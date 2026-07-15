package main

import "fmt"

func sum(numbers ...int) {

	total := 0

	for _, value := range numbers {

		total += value

	}

	fmt.Println(total)

}

func main() {

	sum(10,20)

	sum(10,20,30)

	sum(1,2,3,4,5)

}