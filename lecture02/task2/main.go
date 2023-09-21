package main

import (
	"fmt"
)

// Employee - структура, представляющая сотрудника.
type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

// NewEmployee - функция для создания нового сотрудника.
func NewEmployee(id int, name, position string, salary float64) *Employee {
	return &Employee{ID: id, Name: name, Position: position, Salary: salary}
}

// PrintInfo - метод для вывода информации о сотруднике.
func (e *Employee) PrintInfo() {
	fmt.Printf("ID: %d\n", e.ID)
	fmt.Printf("Имя: %s\n", e.Name)
	fmt.Printf("Должность: %s\n", e.Position)
	fmt.Printf("Зарплата: %.2f\n", e.Salary)
	fmt.Println("-----------------------")
}

// CalculateBonus - метод для расчета бонуса сотрудника.
func (e *Employee) CalculateBonus(percent float64) float64 {
	return e.Salary * percent / 100
}

func main() {
	// Создаем несколько сотрудников.
	employee1 := NewEmployee(1, "Иван Иванов", "Менеджер", 5000.0)
	employee2 := NewEmployee(2, "Анна Петрова", "Разработчик", 3500.0)
	employee3 := NewEmployee(3, "Петр Сидоров", "Дизайнер", 4000.0)

	// Список сотрудников.
	employees := []*Employee{employee1, employee2, employee3}

	// Вывод информации о каждом сотруднике и расчет бонуса.
	for _, employee := range employees {
		employee.PrintInfo()
		bonus := employee.CalculateBonus(10.0)
		fmt.Printf("Бонус: %.2f\n", bonus)
		fmt.Println("-----------------------")
	}
}
