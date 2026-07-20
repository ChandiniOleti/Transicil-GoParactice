package main
import "fmt"
//==empty interface=======

// func display(value interface{}){//here interface means any thing data type we can call any 
// 	fmt.Println(value)
// }
// func main(){
// 	display(100)
// 	display("sdfgtrf")
// 	display(true)
// }

//==creating emp vairables
// func main(){
// 	var value interface{}
// 	value=100
// 	// num:=value.(int)//We need a type assertion.so go can understand the type just value.(type)
// 	// fmt.Println(num+100)
// 	num,ok:=value.(int)
// 	if ok{
// 		fmt.Println(num+30)
// 	}
// }

//========Type switch========

func display(value interface{}){
	switch v:=value.(type){
	case int:
		fmt.Println("Integer: ",v)


	case string:
		fmt.Println("String: ",v)
	default:
		fmt.Println("Unknown type")
	}
}
func main(){
	display(100)
	display("Hello")
	display(true)
}