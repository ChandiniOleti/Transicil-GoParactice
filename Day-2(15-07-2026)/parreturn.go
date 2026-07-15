package main

import "fmt"

/*
==========================================
TYPE 1: No Parameters, No Return Value
==========================================
*/

// func greet() {
// 	fmt.Println("Welcome to Golang!")
// }

/*
==========================================
TYPE 2: Parameters, No Return Value
==========================================
*/

// func add(a int, b int) {
// 	fmt.Println("Sum =", a+b)
// }

/*
==========================================
TYPE 3: No Parameters, With Return Value
==========================================
*/

// func getNumber() int {
// 	return 100
// }

/*
==========================================
TYPE 4: Parameters, With Return Value
==========================================
*/

func multiply(a int, b int) int {
	return a * b
}

func main() {

	// ======================================
	// RUN TYPE 1
	// Uncomment these 2 lines
	// ======================================

	// greet()

	// ======================================
	// RUN TYPE 2
	// Uncomment these 2 lines
	// ======================================

	// add(10, 20)

	// ======================================
	// RUN TYPE 3
	// Uncomment these 3 lines
	// ======================================

	// number := getNumber()
	// fmt.Println(number)

	// ======================================
	// RUN TYPE 4
	// Uncomment these 3 lines
	// ======================================

	result := multiply(10, 20)
	fmt.Println(result)
}