package main
import "fmt"
func main(){
	student := map[string]map[string]string{
	"101":{
		"name": "chanidni",
		"age": "20",
		"city": "guntur",
	},
	"102":{
		"name": "sai",
		"age": "21",
		"city": "hyderabad",
	},
	}
	fmt.Println(student)
}
	