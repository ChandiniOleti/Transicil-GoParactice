package main
import "fmt"
import "time"
func main(){
	go func(){
		fmt.Println("Hello chandini")
	}()
	time.Sleep(time.Second)
	go func(name string){
		fmt.Println("Hello",name)
	}("Oleti")
	time.Sleep(time.Second)
	go func(){
		for i:=1;i<=5;i++{
			fmt.Println(i)
			time.Sleep(500*time.Millisecond)
		}
	}()
	time.Sleep(3*time.Second)
}