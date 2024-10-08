package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

type Manager struct {
	Employee
	Reports []Employee
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "キャスパー",
			ID:   "77",
		},
		Reports: []Employee{},
	}

	var kyasupa Employee = m.Employee
	fmt.Println(kyasupa)
}
