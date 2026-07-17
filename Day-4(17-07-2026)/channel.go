package main
import "fmt"
// func main(){
// 	ch:=make(chan int)//creating the channel
// 	go func(){
// 		ch <-100//sending value
// 	}()
// 	value:= <- ch//receving value
// 	fmt.Println(value)
// }
//sender reciver together
func send(ch chan int){
	ch <-10
	
}
func main(){
	ch:=make(chan int)
	go send(ch)
	
	value:=<-ch
	fmt.Println(value)
	
}