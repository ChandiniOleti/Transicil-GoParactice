package handlers

import (
	"strconv"

	"employeecrudgin/models"

	"github.com/gin-gonic/gin"
)

var Employees = []models.Employee{
	{
		ID:         1,
		Name:       "Chandini",
		Age:        21,
		Department: "IT",
		Salary:     50000,
	},
	{
		ID:         2,
		Name:       "Sai",
		Age:        22,
		Department: "HR",
		Salary:     45000,
	},
}

func GetEmployees(c *gin.Context) {

	c.JSON(200, Employees)

}

func GetEmployeeByID(c *gin.Context) {

	id := c.Param("id")

	empID, err := strconv.Atoi(id)

	if err != nil {

		c.JSON(400, gin.H{
			"error": "Invalid Employee ID",
		})

		return
	}

	for _, employee := range Employees {

		if employee.ID == empID {

			c.JSON(200, employee)
			return
		}

	}

	c.JSON(404, gin.H{
		"error": "Employee Not Found",
	})

}

func CreateEmployee(c *gin.Context) {

	var employee models.Employee

	err := c.BindJSON(&employee)

	if err != nil {

		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})

		return
	}

	Employees = append(Employees, employee)

	c.JSON(201, gin.H{
		"message": "Employee Created Successfully",
		"data": employee,
	})

}

//put method

func UpdateEmployee(c *gin.Context) {

	id := c.Param("id")

	empID, err := strconv.Atoi(id)

	if err != nil {

		c.JSON(400, gin.H{
			"error": "Invalid Employee ID",
		})

		return
	}

	var updatedEmployee models.Employee

	err = c.BindJSON(&updatedEmployee)

	if err != nil {

		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})

		return
	}

	for i, employee := range Employees {

		if employee.ID == empID {

			updatedEmployee.ID = empID

			Employees[i] = updatedEmployee

			c.JSON(200, gin.H{
				"message": "Employee Updated Successfully",
				"data": updatedEmployee,
			})

			return
		}

	}

	c.JSON(404, gin.H{
		"error": "Employee Not Found",
	})

}


//delete operation
func DeleteEmployee(c *gin.Context) {

	id := c.Param("id")

	empID, err := strconv.Atoi(id)

	if err != nil {

		c.JSON(400, gin.H{
			"error": "Invalid Employee ID",
		})

		return
	}

	for i, employee := range Employees {

		if employee.ID == empID {

			Employees = append(Employees[:i], Employees[i+1:]...)

			c.JSON(200, gin.H{
				"message": "Employee Deleted Successfully",
			})

			return
		}

	}

	c.JSON(404, gin.H{
		"error": "Employee Not Found",
	})

}