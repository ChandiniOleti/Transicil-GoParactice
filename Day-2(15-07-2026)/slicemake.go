package main
import "fmt"
func main() {
	numbers := make([]int, 5, 10)
	fmt.Println("Length =", len(numbers))
	fmt.Println("Capacity =", cap(numbers))
}