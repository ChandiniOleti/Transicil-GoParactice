package main
import "fmt"
func main(){
	array:=[]int{10,30,20,40,50}
	max:=array[0]
	for i:=0;i<len(array);i++{
		if array[i]>max{
			max=array[i]
		}
	}
	fmt.Println("Max of array elements:",max)
}