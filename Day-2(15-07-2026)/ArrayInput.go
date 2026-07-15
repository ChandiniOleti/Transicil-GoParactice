package main
import "fmt"
func main(){
	var n [5]int
	fmt.Println("Enter Array elements:")
	for i:=0;i<5;i++{
		fmt.Scan(&n[i])
	}
	fmt.Println("Array elements:",n)
}