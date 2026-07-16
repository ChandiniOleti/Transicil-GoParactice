package main
import "fmt"
import "strconv"
func main(){
	str:="123"
	num,_:=strconv.Atoi(str)
	fmt.Println(num)
	integer:=1234
	stri:=strconv.Itoa(integer)
	fmt.Println(stri)
}