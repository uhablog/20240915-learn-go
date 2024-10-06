package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployee() []Employee {
	newEmployees := []Employee{
		Employee{
			Name: "稲垣祥",
			ID:   "16",
		},
		Employee{
			Name: "キャスパー",
			ID:   "77",
		},
	}

	return newEmployees
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "長谷川健太",
			ID:   "56",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	fmt.Println(m.Name)

	m.Reports = m.FindNewEmployee()
	fmt.Println(m.Employee)
	fmt.Println(m.Reports)
}
