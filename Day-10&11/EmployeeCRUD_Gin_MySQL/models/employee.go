package models

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}