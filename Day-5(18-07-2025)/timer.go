package main
import "fmt"
import "time"
func main(){
	fmt.Println("Waiting")
	// timer:=time.NewTimer(3*time.Second)
	// <-timer.C//it sends the current time through that channel.
	<-time.After(2*time.Second)
	fmt.Println("Timeup!!!")
}