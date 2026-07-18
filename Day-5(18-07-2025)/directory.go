package main
import "fmt"
import "os"
func main(){
	// //==singlr directory=======
	// err:=os.Mkdir("Dir1",0755)
	// if err!=nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Directory is created")
	//====multipple directories=========
	err:=os.MkdirAll("Project/Uploads/Images",0755)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Directories is created")
}