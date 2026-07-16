package main
import "fmt"

type Address struct{
	state string
	city string
}
type student struct{
	name string
	address Address//accesing the struct address 
}
func main(){
	//nested struct
	s:=student{
		name:"Chandini",
		address:Address{
			state:"AP",
			city:"Guntur",
		},
	}
	for _,username:=range student{
	fmt.Println(s)
	}
	//anonymous struct
	ano:=[]struct{
		user string
		gender string
	}{
		{
			user:"Chandini",
			gender:"Female",
		},
		{
			user:"oleti",
			gender:"male",
		},
	}
	fmt.Println(ano)


}

