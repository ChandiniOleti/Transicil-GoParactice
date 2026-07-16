package main
import "fmt"
type student struct{
	name string
	age int
	city string
}
//function with struct parameter
func personname (s student){
	fmt.Println(s)
}
//return struct
func getstudent() student{
	return  student{
		name:"Bhargavi",
		age:30,
		city:"nrt",
	}
}
func main(){
	//functions
	var s student
	s.name="chandini"
	s.age=21
	s.city="Guntur"
	personname(s)
	//pointers
	s2:=student{
		name:"oleti",
		age:40,
		city:"vnk",
	}
	p:=&s2
	fmt.Println(p.name)
	//using *acces
	fmt.Println((*p).age)
	//retrun function
	s3:=getstudent()
	fmt.Println(s3)
}