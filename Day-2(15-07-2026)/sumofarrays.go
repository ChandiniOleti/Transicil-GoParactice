package main
import "fmt"
func main(){
	array:=[]int{10,30,20,40,50}
	sum:=0
	for i:=0;i<len(array);i++{
		sum+=array[i]
	}
	fmt.Println("Sum of array elements:",sum)
	// for _,value:=range array{
	// 	sum+=value
	// }
	// fmt.Println("Sum of array elements using range:",sum)
	
}