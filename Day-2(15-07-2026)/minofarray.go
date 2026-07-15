package main
import "fmt"
func main(){
	arr:=[]int{1,2,3,4,5}
	min:=arr[0]
	for i:=0;i<len(arr);i++{
		if arr[i]<min{
			min=arr[i]
		}
	}
	fmt.Println("Min of array elements:",min)
}