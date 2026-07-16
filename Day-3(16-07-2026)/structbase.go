package main
import "fmt"
type student struct{
	name string
	age int
	city string
}
func main(){
	var s student//object creation
	s.name="Chandini"
	s.age=21
	s.city="Guntur"
	fmt.Println(s)
	//Accessing the fields of the struct
	fmt.Println("Name:",s.name)
	fmt.Println("Age:",s.age)
	fmt.Println("City:",s.city)
	//Struct literals
	s2:=student{
		name:"Oleti",
		age:22,
		city:"Nrt",
	}
	fmt.Println(s2)
}