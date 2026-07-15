package main
import "fmt"
import "math"
func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	temp := num
	var count int
	for temp > 0 {
		temp /= 10
		count++
	}	
	sum := 0
	temp=num
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(count)))
		temp /= 10
	}
	if sum == num {
		fmt.Println("The number is an Armstrong number.")
	} else {
		fmt.Println("The number is not an Armstrong number.")
	}
}