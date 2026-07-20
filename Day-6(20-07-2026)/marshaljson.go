package main
import "fmt"
import "encoding/json"
type student struct{
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
func main(){
	s:=student{
		Id:2001,
		Name:"Chandini",
		Age:21,
	}
	// data,err:=json.Marshal(s)
	data,err:=json.MarshalIndent(s,""," ")//want (any, string, string)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}