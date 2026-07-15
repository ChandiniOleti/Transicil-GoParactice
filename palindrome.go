package main
import "fmt"
func main(){
	var num int
	var temp int
	var rem int
	var rev int=0
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	temp = num
	i:=1
	for i<=num{
		rem=num%10
		rev=rev*10+rem
		num=num/10
	}
	if temp==rev{
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println("The number is not a palindrome.")
	}
}