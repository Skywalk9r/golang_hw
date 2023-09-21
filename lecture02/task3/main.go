package main

import (
	"fmt"
)

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

func NewEmployee(id int, name, position string, salary float64) *Employee {
	return &Employee{ID: id, Name: name, Position: position, Salary: salary}
}

func (e *Employee) PrintInfo() {
	fmt.Printf("ID: %d\n", e.ID)
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Job title: %s\n", e.Position)
	fmt.Printf("Salary: %.2f\n", e.Salary)
	fmt.Println("-----------------------")
}

func (e *Employee) CalculateBonus(percent float64) float64 {
	return e.Salary * percent / 100
}

func main() {
	employee1 := NewEmployee(1, "Aizhan Tulendinova", "Manager", 5000.0)
	employee2 := NewEmployee(2, "Nurbol Seitzhanov", "Developer", 3500.0)
	employee3 := NewEmployee(3, "Madiyar Lashbayev", "Designer", 4000.0)

	employees := []*Employee{employee1, employee2, employee3}

	for _, employee := range employees {
		employee.PrintInfo()
		bonus := employee.CalculateBonus(10.0)
		fmt.Printf("Bonus: %.2f\n", bonus)
		fmt.Println("-----------------------")
	}
}
