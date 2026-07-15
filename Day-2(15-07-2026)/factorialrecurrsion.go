package main
import "fmt"
func factorial(n int) int{
	if n==0{
		return 1
	}
	return factorial(n-1) * n
}
func main(){
	var result int
	fmt.Print("Enter the number: ")
	fmt.Scan(&result)
	fmt.Println(factorial(result))
}