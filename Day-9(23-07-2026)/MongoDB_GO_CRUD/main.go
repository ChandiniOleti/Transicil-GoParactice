package main

import (
	"fmt"

	"mongodbgocrud/crud"
	"mongodbgocrud/database"
)

func main() {

	client, err := database.ConnectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer client.Disconnect(nil)

	fmt.Println("MongoDB Connected Successfully")

	crud.InsertEmployee(client)
	crud.GetAllEmployees(client)
	crud.GetEmployeeByName(client, "Chandini")
	crud.UpdateEmployeeSalary(client, "Chandini", 65000)
	crud.DeleteEmployee(client, "Chandini")
}