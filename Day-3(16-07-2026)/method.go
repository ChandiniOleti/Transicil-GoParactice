package main
import "fmt"
type student struct{
	name string
	age int
}
func (s student) Display(){
	fmt.Println("Name: ",s.name)
	fmt.Println("Age: ",s.age)
}
func main(){
	stud:=student{
		name:"chandini",
		age:21,
	}
	stud.Display()
}