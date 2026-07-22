package main

import (
	"fmt"

	"employeecrud/crud"
	"employeecrud/database"
	"employeecrud/models"
)

func main() {

	// Connect to Database
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	// Create Employee Object
	// emp := models.Employee{
	// 	Name:       "Chandini",
	// 	Age:        21,
	// 	Department: "IT",
	// 	Salary:     50000,
	// 	City:       "Guntur",
	// }
	//if we need from the user we need to create varable and and use sacn method totake input
	var emp models.Employee

	fmt.Print("Enter Name: ")
	fmt.Scan(&emp.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&emp.Age)

	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter Department: ")
	//department, _ := reader.ReadString('\n')==========To read the entire line, use bufio.NewReader.

	fmt.Print("Enter Department: ")
	fmt.Scan(&emp.Department)//fmt.Scanln() → Reads input until Enter.

	fmt.Print("Enter Salary: ")
	fmt.Scan(&emp.Salary)

	fmt.Print("Enter City: ")
	fmt.Scan(&emp.City)

	// Insert Employee
	err = crud.InsertEmployee(db, emp)
	if err != nil {
		fmt.Println("Insert Error:", err)
		return
	}

	fmt.Println("Employee Inserted Successfully")

	employees, err := crud.GetEmployees(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, emp := range employees {

		fmt.Println("ID:", emp.ID)
		fmt.Println("Name:", emp.Name)
		fmt.Println("Age:", emp.Age)
		fmt.Println("Department:", emp.Department)
		fmt.Println("Salary:", emp.Salary)
		fmt.Println("City:", emp.City)
		fmt.Println("-------------------------")
	}
}