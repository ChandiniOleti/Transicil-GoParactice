package main
import "fmt"
import "os"
func main(){
	//=======directory rename============
// 	err:=os.Rename("Dir1","Practice")
// if err!=nil{
// 		fmt.Println(err)
// 		return
// 	}
	// fmt.Println("Directory rreanamed")
	


	//===directory removed=========
	rem:=os.Remove("Dict")
	if rem!=nil{
		fmt.Println(rem)
		return
	}
	fmt.Println("Directory is removed")
}