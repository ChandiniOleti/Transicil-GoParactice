package crud

import (
	"database/sql"
	"employeecrud/models"
)

func GetEmployees(db *sql.DB) ([]models.Employee, error) {

	query := "SELECT * FROM employeego"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {

		var emp models.Employee

		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Age,
			&emp.Department,
			&emp.Salary,
			&emp.City,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, emp)
	}

	return employees, nil
}