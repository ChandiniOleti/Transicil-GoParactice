package main
import "fmt"
//send-only channel syntax chan<-
// func send(ch chan<- int){
// 	ch<-50
// }
// func main(){
// 	ch:=make(chan int)
// 	go send(ch)
// 	fmt.Println(<-ch)
// }


//reciveonly channel syntax <-chan
func recv(ch <-chan int){
	value:=<-ch
	fmt.Println(value)
}
func main(){
	ch:=make(chan int)
	go func(){
		ch<-200
	}()
	recv(ch)
}