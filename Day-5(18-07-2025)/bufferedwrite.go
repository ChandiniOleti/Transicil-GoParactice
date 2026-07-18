package main
import "fmt"
import "os"
import "bufio"
func main(){
	file,err:=os.Create("sample.txt")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer:=bufio.NewWriter(file)
	writer.WriteString("Hello\n")
	writer.WriteString("Chandini\n")
	writer.Flush()
	fmt.Println("Writed successfully")
}
