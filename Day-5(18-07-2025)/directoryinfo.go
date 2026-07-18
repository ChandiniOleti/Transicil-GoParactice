package main
import "fmt"
import "os"
func main(){

info, err := os.Stat("Practice")

if err != nil {
	fmt.Println(err)
	return
}

fmt.Println(info.Name())
fmt.Println(info.Size())
fmt.Println(info.IsDir())
fmt.Println(info.Mode())
fmt.Println(info.ModTime())
}