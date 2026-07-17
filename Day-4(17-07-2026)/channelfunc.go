package main
import "fmt"
func square(num int ,ch chan int){
	ch<-num*num

}
func main(){
	ch:=make(chan int)
	go square(5,ch)
	value:=<-ch
	fmt.Println(value)
}