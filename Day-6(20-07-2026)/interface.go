package main
import "fmt"
type Animal interface{
	Speak()
}
type Dog struct{}
func (d Dog) Speak() {
	fmt.Println("Bow Bow..........")
}
type Cat struct{}
func (c Cat) Speak() {
	fmt.Println("Meow Meow........")
}
func main(){
	var a Animal
	a=Dog{}
	a.Speak()
	a=Cat{}
	a.Speak()
}