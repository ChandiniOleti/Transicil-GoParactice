package main

import "fmt"
func main(){
	student:=make(map[string]string)
	student["name"]="chanidni"
	student["age"]="20"
	student["city"]="guntur"
	fmt.Println(student)
	for key,value:=range student{
		fmt.Println(key,value)
	}
	//add new key value pair
	student["branch"] = "CSE"
	fmt.Println(student)
	//delete key value pair
	delete(student,"age")
	fmt.Println(student)
	//update key value pair
	student["city"] = "Hyderabad"
	fmt.Println(student)
	//key exixts or not
	value,exists:=student["name"]
	fmt.Println(value)
	fmt.Println(exists)
}