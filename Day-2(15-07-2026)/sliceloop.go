package main
import "fmt"
func main(){

	var arr [5]int
	fmt.Println("Enter array elements:")
	for i:=0;i<5;i++{
		fmt.Scan(&arr[i])

	}
	fmt.Println("Array elements:", arr)
	//====================range========
	// for index,value:=range arr{
	// 	fmt.Println(index,value)
	// }
	//===================
	// Bakend example
	// users := []string{

	// "Rahul",

	// "Chandini",

	// "Ravi",

	// }
	// for _,user:=range users{

	// fmt.Println(user)

	// }
}
