package main
import "fmt"
import "regexp"
import "bufio"
import "os"
func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the sentence:")
	text, _ := reader.ReadString('\n')
	word:=regexp.MustCompile("Go")
	result:=word.FindAllString(text,-1)
	fmt.Println("Words",result)
	fmt.Println("Count",len(result))
}