package main
import "fmt"
func named(a,b int) (result int) {
	result = a + b
	return
}
func main(){
	fmt.Println(named(10,20))
}