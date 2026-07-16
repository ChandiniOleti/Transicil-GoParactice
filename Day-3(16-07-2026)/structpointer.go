package main
import "fmt"
type Student struct {
	Name string
	Age  int
}
func main() {
	s := Student{
		Name: "Chandini",
		Age:  21,
	}
	p := &s
	fmt.Println(p.Name)
	p.Age=20
	fmt.Println(s.Age)
}