package main
import "fmt"
import "errors"
func validateage(age int) error{
	if age<18{
		return errors.New("Age must be greater than 18")//fixed message
	}
	return nil
}
func main(){
	var num int
	fmt.Print("Enter the age: ")
	fmt.Scan(&num)
	err:=validateage(num)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Eligible")
}