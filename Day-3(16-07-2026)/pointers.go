package main
import "fmt"
func user(num int){
	num=100
}
func student(n *int){
	*n=1000
}
func main(){
	x:=10
	user(x)
	fmt.Println("pass by value: ",x)
	y:=10
	student(&y)
	fmt.Println("pass by reference: ",y)
}