package main
import "fmt"
type student struct {
	name string
	age int
}
func main(){
	//struct using slice
	s:=[2]student{
		{name: "Chandini", age: 20},
		{name: "Oleti", age: 22},
	}
	//using range
	for _,stud:=range s{
	fmt.Println(stud)
	}
	//sruct using slice
	s2:=[]student{
		{name:"Chandini",age:21},
		{name:"Bhargavi",age:19},
		{name:"Oleti",age:30},
	}
	for _,stu:=range s2{
		fmt.Println(stu)
	}

}