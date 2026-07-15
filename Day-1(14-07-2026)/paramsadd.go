package main
import "fmt"
// func add(a,b int) int{
	
// 	return a+b
// }
func add(a int, b int) (int, int){
	sum:=a+b
	product:=a*b
	return sum,product
}
func main(){
	var a,b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)
	sum, product := add(a, b)
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
}