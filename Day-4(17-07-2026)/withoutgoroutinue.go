package main

import (
	"fmt"
	"time"
)

func task1() {
	fmt.Println("Task 1 Started")
	time.Sleep(2 * time.Second)
	fmt.Println("Task 1 Completed")
}

func task2() {
	fmt.Println("Task 2 Started")
	time.Sleep(2 * time.Second)
	fmt.Println("Task 2 Completed")
}

func main() {

	task1()
	task2()

}