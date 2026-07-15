package main
import "fmt"
func main(){
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	for num!=0{
		rem:=num%10
		rev:=rev*10+rem
		num=num/10
	
	}
	fmt.Println("Reversed number: ",rev)
}