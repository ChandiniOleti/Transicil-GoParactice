package main
import "fmt"
import "sync"
func main(){
	var wg sync.WaitGroup
	ch:=make(chan int)
	wg.Add(1)
	go func(){
		
		ch<-100
		defer wg.Done()
	}()
	value:=<-ch
	wg.Wait()
	fmt.Println(value)

}