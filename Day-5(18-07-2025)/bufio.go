package main
import "fmt"
import "os"
import "bufio"
import "io"
func main(){
	// //===ReadString======
	// file,err:=os.Open("student.txt")
	// if err!=nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()
	// reader:=bufio.NewReader(file)
	// for {
	// 	line,err:=reader.ReadString	('\n')
	// 	fmt.Println(line)
	// 	if err==io.EOF{//Finished Reading
	// 		break
	// 	}
	// 	if err!=nil{
	// 		fmt.Println(err)
	// 		return
	// 	}
		
		
	// }
	//======ready charecter by charecter=======
	file,_:=os.Open("student.txt")
	defer file.Close()
	reader:=bufio.NewReader(file)
	for {
		char,_,err:=reader.ReadRune()
		
		if err==io.EOF{
			break
		}
		fmt.Printf("%c\n",char)
	}
}