package main
import "fmt"
func main() {
	a:=0
	b:=1
	var n int
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&n)
	fmt.Print("Fibbnocci series:")
	for i:=1;i<=n;i++{
		fmt.Print(a," ")
		c:=a+b
		a=b
		b=c

	}
	
}