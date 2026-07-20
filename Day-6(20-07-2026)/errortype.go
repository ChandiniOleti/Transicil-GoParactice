package main
import "fmt"
import "os"
import"errors"
func main(){
	file,err:=os.Open("chandini.txt")
	// if err!=nil{
	// 	fmt.Println("Error no such file",err)
	// 	return
	// }
	// defer file.Close()
	// fmt.Println(file)

	//Checking Specific Errors
	if errors.Is(err,os.ErrNotExist){
		fmt.Println("File not found")
	}
	fmt.Println(file)
}