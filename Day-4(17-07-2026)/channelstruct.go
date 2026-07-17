package main
import "fmt"
type Student struct{
	name string
	age int
}
func main(){
	ch:=make(chan Student)
	go func(){
		ch<-Student{"chandini",21}
		ch<-Student{"oleti",22}
	}()
	student1:=<-ch
	student2:=<-ch
	fmt.Println(student1)
	fmt.Println(student2)
}