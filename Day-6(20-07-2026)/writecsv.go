package main

import (
	"encoding/csv"
	"os"
)
func main() {
	file, _ := os.Create("students.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"101", "Chandini", "21"})
	writer.Write([]string{"102", "Sai", "22"})
}