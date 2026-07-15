package main

import "fmt"
import "strings"

func main() {

	name := "chandini"

	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToLower(name))
	email := "chandini@gmail.com"
	if strings.Contains(email, "@") {
		fmt.Println("Valid Email")
	}
	url := "https://google.com"
	if strings.HasPrefix(url, "https") {
		fmt.Println("Secure URL")
	}
	if strings.HasSuffix(url, ".com") {
		fmt.Println("Commercial URL")
	}
	text := "my name is oleti"
	fmt.Println(strings.Replace(text, "oleti", "chandini", 1))
	text1 := "Apple,Banana,Mango"
	fmt.Println(strings.Split(text1, ","))
	fruits := []string{"Apple", "Banana", "Mango"}
	fmt.Println(strings.Join(fruits, "-"))


}


