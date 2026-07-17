package main
import "fmt"
import "sync"
func hello(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Hello Chandini")
}
func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg)
	wg.Wait()
}