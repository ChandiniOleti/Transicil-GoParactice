package main
import "fmt"
type Student struct{
name string
age int
}
//Value Receiver
func (s Student) Updatename(){
	s.name="Oleti"
}
//Pointer Receiver
func (s *Student) Updateage(){
	s.age=20
}
func main(){
	s1:=Student{
		name:"Chandini",
	}
	s1.Updatename()
	fmt.Println(s1.name)
	s2:=Student{
		age:30,
	}
	s2.Updateage()
	fmt.Println(s2.age)
}