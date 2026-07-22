package main

import (
	"context"
	"fmt"

	"employeecrudsqlc/database"
	db "employeecrudsqlc/internal/db"
)

func main() {

	// Connect to Database
	conn, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Create SQLC Queries Object
	queries := db.New(conn)

	// Context is required by SQLC
	ctx := context.Background()

	// params := db.CreateEmployeeParams{

	// 	Name:       "Chandini",
	// 	Age:        21,
	// 	Department: "IT",
	// 	Salary:     50000,
	// 	City:       "Guntur",
	// }

	// err = queries.CreateEmployee(ctx, params)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Employee Inserted Successfully")


	//======reading the all values

	employees, err := queries.ListEmployees(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, emp := range employees {

		fmt.Println(emp.ID)
		fmt.Println(emp.Name)
		fmt.Println(emp.Age)
		fmt.Println(emp.Department)
		fmt.Println(emp.Salary)
		fmt.Println(emp.City)
		fmt.Println("----------------")
	}


	//reading only 1 row
	employee, err := queries.GetEmployee(ctx, 1)//here 1 isthe param

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(employee)


	//---delete the row with param
	err1 := queries.DeleteEmployee(ctx, 2)

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println("Deleted Successfully")
}