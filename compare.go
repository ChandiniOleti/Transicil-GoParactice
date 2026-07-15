
import "fmt"

func main() {
	a := 12
	b := 10
	if a == b {
		fmt.Println("Both numbers are equal")

	} else if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("b is greater than a")
	}

}