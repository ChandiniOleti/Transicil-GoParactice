package main
import "fmt"
import "regexp"
func main(){
	csv:="101,chandini,Guntur,21"
	words:=regexp.MustCompile(",")
	data:=words.Split(csv,-1)
	for _,result:=range data{
		fmt.Println(result)
	}
}