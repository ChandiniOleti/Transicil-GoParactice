package main

import "fmt"
import "os"
import "encoding/json"
import "sort"
//creating the struct
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
	err = json.Unmarshal(data, &s)//Convert JSON into Go Structs
	if err!=nil{
		fmt.Println(err)
		return
	}
	// //==search by id
	// id:=102
	// found :=false
	// for _,stud:=range s{
	// 	if stud.ID==id{
	// 		fmt.Println(stud)
	// 		found=true
	// 		break
	// 	}
	// }
	// if !found{
	// 	fmt.Println("Student id is not found")
	// }

	// //===search by name
	// name:="Chandini"
	// found=false
	// for _,stun:=range s{
	// 	if stun.Name==name{
	// 		fmt.Println(stun)
	// 		found=true
	// 		break
	// 	}
	// }
	// if !found{
	// 	fmt.Println("Student name not  found")
	// }

	// //========filter by age====
	// for _,stuage:=range s{
	// 	if stuage.Age>20{
	// 		fmt.Println(stuage)
	// 	}
	// }

	// //========filter by city
	// for _,stucity:=range s{
	// 	if stucity.City=="Guntur"{
	// 		fmt.Println(stucity)
	// 	}
	// }


	//=======sort by Age Ascending
	
	sort.Slice(s, func(i, j int) bool {
	return s[i].Age < s[j].Age
	})
	for _,stu:=range s{
		fmt.Println(stu)
	}

	//=====sort by age Desending======
	sort.Slice(s, func(i, j int) bool {
	return s[i].Age > s[j].Age
	})
	for _,stu:=range s{
		fmt.Println(stu)
	}




	//=====add new student
	newstudent:=student{
		ID:104,
		Name:"Oleti",
		Age:19,
		City:"Nrt",
	}
	s=append(s,newstudent)
	fmt.Println(s)//still the file is nor overwrited we need convert this into the marshal than go to json

	//=====update the student city
	for i:=range s{
		if s[i].ID==102{
			s[i].City="Bengalore"
			fmt.Println(s)//here is updated the city name but not in the main file
		}
	}
	//just upadting the content to the in the go to json formate
	jsondata,err:=json.MarshalIndent(s,""," ")//Save Updated Data Back to JSON
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsondata))

	//=========Write Back to File

	err=os.WriteFile("students.json",jsondata,0644)
	if err!=nil{
		fmt.Println(err)
		return
	}

}