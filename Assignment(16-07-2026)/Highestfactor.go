package main
import "fmt"
func main(){
	var n,k int
	arr:=[]int{}
	fmt.Println("Enter the N value")
	fmt.Scan(&n)
	fmt.Println("Enter the k value")
	fmt.Scan(&k)
	for i:=1;i<=n;i++{
		if n%i==0{
			arr=append(arr,i)			
		}		
	}
	if k>len(arr){
		fmt.Println(1)
	}else{
	fmt.Println("kth highest factor of N: "arr[k])
	}
}