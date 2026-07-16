package main
import "fmt"
type calculator struct{}
func (c calculator) Add(a,b int){
	fmt.Println(a+b)
}
//return type
func (c calculator) sub(a,b int) int{
	return a-b
}
func (c calculator) mul(a,b int) int{
	return a*b
}
func (c calculator) div(a,b int) int{
	return a/b
}
func main(){
	cal:=calculator{}
	var n1,n2 int
	fmt.Print("enter number1: ")
	fmt.Scan(&n1)
	fmt.Print("enter number2: ")
	fmt.Scan(&n2)
	cal.Add(n1,n2)
	subract:=cal.sub(n1,n2)
	fmt.Println(subract)
	multiple:=cal.mul(n1,n2)
	fmt.Println(multiple)
	division:=cal.div(n1,n2)
	fmt.Println(division)
}