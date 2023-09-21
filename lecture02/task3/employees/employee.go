package employees

import "fmt"

type Employee interface {
	GetName() string
	GetSalary() float64
	Work() string
}

type CommonFields struct {
	Name   string
	Salary float64
}

func (c CommonFields) PrintInfo() {
	fmt.Printf("Имя: %s\n", c.Name)
	fmt.Printf("Зарплата: %.2f\n", c.Salary)
}
