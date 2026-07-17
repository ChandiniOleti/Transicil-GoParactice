package main
import "fmt"
import "time"
//======multiplexing multiple channels=====
// func main(){
// 	ch1:=make(chan int)
// 	ch2:=make(chan string)
// 	go func(){
// 		ch1<-21
// 	}()
// 	go func(){
// 		ch2<-"chandini"
// 	}()
// 	select{
// 	case msg1:=<-ch1:
// 		fmt.Println(msg1)
	
// 	case msg2:=<-ch2:
// 		fmt.Println(msg2)
// 	}
// }


//======select with Time Delay
func main(){
	ch1:=make(chan int)
	ch2:=make(chan string)
	go func(){
		ch1<-21
	}()
	go func(){
		time.Sleep(2*time.Second)
		ch2<-"chandini"
	}()
	select{
	case msg1:=<-ch1:
		fmt.Println(msg1)
	
	case msg2:=<-ch2:
		fmt.Println(msg2)
	
	//Timeout using time.After()
	case <-time.After(4*time.Second):
		fmt.Println("Timeout!")
	}
}