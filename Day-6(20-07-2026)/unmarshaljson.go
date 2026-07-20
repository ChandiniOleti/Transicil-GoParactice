package main
import "fmt"
import "encoding/json"
type student struct{
	Id int
	Name string
	Age int
}
func main(){
	jsondata:=`{"Id":101,"Name":"Chandini","Age":21}`
	var s student
	err:=json.Unmarshal([]byte(jsondata),&s)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}