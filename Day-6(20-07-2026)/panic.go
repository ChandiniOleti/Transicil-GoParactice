package main
import "fmt"
//===panicmessage======
// func main(){
// 	defer fmt.Println("Deffered function")
// 	fmt.Println("Program started")
// 	panic("Unexpected error")
// 	fmt.Println("Program end")
// }



//=====stack unwinding============

// func demo(){
// 	defer fmt.Println("demo defer")
// 	panic("Error")
// }
// func main(){
// 	defer fmt.Println("Main defer")
// 	demo()
// }




//===========recover===========
//recover() catches a panic and prevents the program from crashing.
//It works only inside a deferred function.
func main(){
	defer func(){
		message:=recover()
		fmt.Println("Recoverd: ",message)
		fmt.Println("Hello")
	}()
	panic("Server Crashed")
	fmt.Println("Hello Chandini")
}