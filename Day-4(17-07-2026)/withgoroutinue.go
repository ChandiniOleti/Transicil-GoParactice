package main

import (
	"fmt"
	"time"
)

func printnumbers(){
	for i:=1;i<=5;i++{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {

	go printnumbers()
	time.Sleep(6*time.Second)

}