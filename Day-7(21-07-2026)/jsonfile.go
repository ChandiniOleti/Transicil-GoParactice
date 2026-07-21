package main

import "fmt"
import "os"
import "encoding/json"
type student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}
func main() {
	data, err := os.ReadFile("students.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var s []student
	err = json.Unmarshal(data, &s)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}