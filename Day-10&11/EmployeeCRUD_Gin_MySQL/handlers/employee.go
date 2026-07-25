package handlers

import (
	"net/http"

	"employeecrud/config"
	"employeecrud/models"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {

	var employee models.Employee

	// Read JSON body
	err := c.ShouldBindJSON(&employee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


	// Insert into database
	query := `
	INSERT INTO employees(name, age, department, salary)
	VALUES (?, ?, ?, ?)
	`

	_, err = config.DB.Exec(
		query,
		employee.Name,
		employee.Age,
		employee.Department,
		employee.Salary,
	)


	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(201, gin.H{
		"message": "Employee created successfully",
		"employee": employee,
	})
}

//get method

func GetEmployees(c *gin.Context) {

	rows, err := config.DB.Query(
		"SELECT id, name, age, department, salary FROM employees",
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()


	var employees []models.Employee


	for rows.Next() {

		var employee models.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Age,
			&employee.Department,
			&employee.Salary,
		)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}


		employees = append(employees, employee)
	}


	c.JSON(200, employees)

}

//get by the id

func GetEmployeeByID(c *gin.Context) {

	id := c.Param("id")

	var employee models.Employee


	err := config.DB.QueryRow(
		"SELECT id, name, age, department, salary FROM employees WHERE id=?",
		id,
	).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Age,
		&employee.Department,
		&employee.Salary,
	)


	if err != nil {

		c.JSON(404, gin.H{
			"message": "Employee not found",
		})

		return
	}


	c.JSON(200, employee)

}

func UpdateEmployee(c *gin.Context) {

	id := c.Param("id")

	var employee models.Employee


	// Read JSON body
	err := c.ShouldBindJSON(&employee)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}


	// Update database
	query := `
	UPDATE employees
	SET name=?, age=?, department=?, salary=?
	WHERE id=?
	`


	_, err = config.DB.Exec(
		query,
		employee.Name,
		employee.Age,
		employee.Department,
		employee.Salary,
		id,
	)


	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(200, gin.H{
		"message": "Employee updated successfully",
	})

}

//delete using id
func DeleteEmployee(c *gin.Context) {

	id := c.Param("id")


	_, err := config.DB.Exec(
		"DELETE FROM employees WHERE id=?",
		id,
	)


	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(200, gin.H{
		"message": "Employee deleted successfully",
	})

}