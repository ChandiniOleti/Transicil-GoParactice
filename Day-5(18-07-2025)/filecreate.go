package main
import "fmt"
import "os"
// import "bufio"
func main(){
	//creating the file
	file,err :=os.Create("exp1.txt")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	//writing the file
	file.WriteString("Hello")
	file.WriteString("Chandini welcome")
	file.WriteString("Learn Go")
	

	// // fmt.Println("File is created succesfully")
	// //============reading the file============
	// data,err:=os.ReadFile("exp1.txt")
	// if err!=nil{
	// 	fmt.Println(err)
	// 	return
	// }
	
	// fmt.Println("Data is readed succesfully",string(data))

	// //========Opening the file============
	// file,err:=os.Open("exp1.txt")
	// if err!=nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()
	// scanner:=bufio.NewScanner(file)
	// for scanner.Scan(){
	// 	fmt.Println(scanner.Text())
	// }
	// if err:=scanner.Err();err!=nil{
	// 	fmt.Println(err)
	// }
}