package employees

type Worker struct {
	CommonFields
	Worksite string
}

func (w Worker) GetName() string {
	return w.Name
}

func (w Worker) GetSalary() float64 {
	return w.Salary
}

func (w Worker) Work() string {
	return "Выполнение работ на " + w.Worksite
}
