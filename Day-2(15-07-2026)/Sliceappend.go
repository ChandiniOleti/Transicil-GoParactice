package main
import "fmt"
func main(){
	n:=[]int{1,2,3,4,5}
	n=append(n,6,7,8)
	fmt.Println(n)
}