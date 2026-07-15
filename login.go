package main

import "fmt"

func main() {
	var email, password string
	fmt.Print("Email:")
	fmt.Scan(&email)
	fmt.Print("Password:")
	fmt.Scan(&password)
	if email == "abc@gmail.com" {
		if password == "123456" {
			fmt.Println("Login Successful")
		} else {
			fmt.Println("Wrong Password")
		}
	} else {
		fmt.Println("Email Not Found")
	}

}
