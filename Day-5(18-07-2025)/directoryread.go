package main
import "fmt"
import "os"
func main(){
	// files,err:=os.ReadDir("Dir1")
	// if err!=nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// for _,file:=range files{
	// 	fmt.Println(file.Name())
	// }
//}

	//====Current directory=======
	path, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(path)


}