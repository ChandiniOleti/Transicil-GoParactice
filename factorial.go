package main
import "fmt"
func main(){
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	fact := 1
	for num>=1{
		fact = fact * num
		num = num - 1
	}
	fmt.Println("Factorial: ", fact)
}