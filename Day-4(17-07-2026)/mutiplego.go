package main

import "fmt"
import	"time"


func task(name string) {

	for i := 1; i <= 3; i++ {

		fmt.Println(name, i)

		time.Sleep(time.Second)

	}

}

func main() {

	go task("A")

	go task("B")

	time.Sleep(4 * time.Second)

}