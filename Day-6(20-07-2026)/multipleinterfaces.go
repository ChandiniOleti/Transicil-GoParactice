package main

import "fmt"

type Vehicle interface {
	Start()
	Stop()
}

type Car struct{}

func (c Car) Start() {
	fmt.Println("Car Started")
}

func (c Car) Stop() {
	fmt.Println("Car Stopped")
}

func main() {

	var v Vehicle

	v = Car{}

	v.Start()
	v.Stop()
}