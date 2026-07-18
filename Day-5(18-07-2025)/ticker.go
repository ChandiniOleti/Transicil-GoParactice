package main
import "fmt"
import "time"
func main(){
	ticker:=time.NewTicker(1*time.Second)
	// for i:=1;i<=5;i++{
	// 	<-ticker.C
	// 	fmt.Println("Tick",i)
	// }
	// ticker.Stop()//it is very importtant to stop the ticker otherwise Ticker keeps running forever.
	count:=0
	for range ticker.C{
		count++
		fmt.Println("Tick",count)
		if count==5{
			ticker.Stop()
			break	
		}
	}
}