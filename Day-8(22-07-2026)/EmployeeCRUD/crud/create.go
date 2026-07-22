package crud

import (
	"database/sql"
	"employeecrud/models"
)

func InsertEmployee(db *sql.DB, emp models.Employee) error {
	//The ? are placeholders. Go replaces them safely with the values you pass.
	query := `INSERT INTO employeego(name,age,department,salary,city)
	          VALUES(?,?,?,?,?)`//db → Database connection emp → Employee data from type struct 

	_, err := db.Exec(
		query,
		emp.Name,
		emp.Age,
		emp.Department,
		emp.Salary,
		emp.City,
	)//Exec() executes queries that don't return rows, such as: insert ,delete,update========The _ ignores the sql.Result because we don't need it here.

	return err
}