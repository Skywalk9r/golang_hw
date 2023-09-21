package main

import (
	"fmt"
	"golang_lectures/lecture02/task3/employees"
)

func main() {
	manager := employees.Manager{
		CommonFields: employees.CommonFields{Name: "Иван Петрович", Salary: 5000.0},
		Department:   "Производство",
	}

	worker := employees.Worker{
		CommonFields: employees.CommonFields{Name: "Анна Ивановна", Salary: 3000.0},
		Worksite:     "Стройка",
	}

	employeeList := []employees.Employee{manager, worker}

	for _, employee := range employeeList {
		employee.PrintInfo()
		fmt.Println("Обязанности:", employee.Work())
		fmt.Println("-----------------------")
	}
}
