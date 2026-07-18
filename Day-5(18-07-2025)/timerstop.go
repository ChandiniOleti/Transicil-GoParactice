package main

import "fmt"
import "time"

func main() {

	timer := time.NewTimer(5 * time.Second)
	
	time.Sleep(2 * time.Second)
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	<-timer.C
	// timer.Stop()//you don't need it anymore.
	// timer.Reset(3*time.Second)//Reuse the same timer.

	fmt.Println("Timer Stopped")

}