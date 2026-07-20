package main
import "fmt"
import "encoding/csv"
import "os"
func main(){
	file,err:=os.Open("student.csv")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader:=csv.NewReader(file)
	records, err := reader.ReadAll()

	if err!=nil{
		fmt.Println(err)
		return
	}
	for _,data:=range records{
		fmt.Println("ID: ",data[0])//for accesing single value
		fmt.Println("Name: ",data[1])
		fmt.Println("Age: ",data[2])
		fmt.Println(data)
	}
}