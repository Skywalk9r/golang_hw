package employees

type Manager struct {
	CommonFields
	Department string
}

func (m Manager) GetName() string {
	return m.Name
}

func (m Manager) GetSalary() float64 {
	return m.Salary
}

func (m Manager) Work() string {
	return "Управление отделом " + m.Department
}
