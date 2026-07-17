package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for {//here with out for it prints only one value so for multiple values to print we using for
		select {
		case value, ok := <-ch:
			if !ok {
				fmt.Println("Channel Closed")
				return
			}

			fmt.Println(value)

		}

	}

}