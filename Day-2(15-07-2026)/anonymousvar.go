package main
import "fmt"
func main(){
	greet:=func(){
		fmt.Println("Welcome to Golang!")
	}
	greet()
}