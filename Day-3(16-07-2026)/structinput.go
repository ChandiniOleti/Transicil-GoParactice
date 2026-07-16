package main
import "fmt"
type student struct{
	name string
	age int

}
func main(){
	var s student
	fmt.Print("Enter the name: ")
	fmt.Scan(&s.name)
	fmt.Print("Enter the age: ")
	fmt.Scan(&s.age)
	fmt.Println(s)
}