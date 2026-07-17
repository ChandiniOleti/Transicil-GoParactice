package main
import "fmt"
// func main(){//without the go routinue using capacity
// 	ch:=make(chan int,3)
// 	ch<-10
// 	ch<-20
// 	stud1:=<-ch
// 	stud2:=<-ch
// 	fmt.Println(stud1)
// 	fmt.Println(stud2)
// 	fmt.Println(len(ch))//to get the lenght don't print it first if so it gives the 0
// 	fmt.Println(cap(ch))
// }
//buffered with goroutinues
func main(){
	ch:=make(chan int,5)
	go func(){
		ch<-20
		ch<-30
		ch<-40

	}()
	value1:=<-ch
	value2:=<-ch
	value3:=<-ch
	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3)
}