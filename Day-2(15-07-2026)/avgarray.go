package main
import "fmt"
func main(){
	array:=[]int{10,20,30,40,50}
	sum:=0
	for i:=0;i<len(array);i++{
		sum+=array[i]
	}
	avg:=float64(sum)/float64(len(array))
	fmt.Println("Average of array elements:",avg)
}