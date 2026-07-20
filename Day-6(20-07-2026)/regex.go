package main
import "fmt"
import "regexp"
func main(){
	//======matching strings=======
	// match,err:=regexp.MatchString("Go","Learn Go Language")
	

	//=======complie======Regex patterns are first converted into an internal format.====
	// comp, err := regexp.Compile("Go")

	//======mustcomplie doesnot need err to return

	// msc := regexp.MustCompile("Go")
	// if err!=nil{
	// 	fmt.Println(err)
	// 	return
	// }

	//===FindString only the first string
	fs := regexp.MustCompile("Go")
	// fmt.Println(fs.FindString("Go Learning Go"))
	

	//=======findallstrings  we keep -1 which means it sholud return all go match word in string if we need only 2 we sholud mention only 2
	// fmt.Println(fs.FindAllString("Go Java Go Python Go", -1))


	//==replace all matching strings
	result:=fs.ReplaceAllString(
		"Go is fast",
		"GOlang",
	)
	fmt.Println(result)
}