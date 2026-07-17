package main

import (
	"fmt"
	"sync"
)

func print(name string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(name)

}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go print("Go", &wg)

	go print("Java", &wg)

	go print("Python", &wg)

	wg.Wait()

	fmt.Println("Finished")

}