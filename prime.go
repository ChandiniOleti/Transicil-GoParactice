package main
import "fmt"
func main(){
	var num int
	isPrime := true
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	for i:=2;i<num;i++{
		if num%i==0{
			isPrime=false
			break
		}else{
			isPrime=true
		}
	}
	if isPrime {
		fmt.Println("The number is a prime number.")
	} else {
		fmt.Println("The number is not a prime number.")
	}
}