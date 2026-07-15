package main
import "fmt"
func main(){
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	sum:=0
	for i:=1;i<=num;i++{
		sum+=i
	}
	fmt.Println("Sum of the digits : ",sum)
}