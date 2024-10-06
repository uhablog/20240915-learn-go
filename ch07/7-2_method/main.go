package main

import "fmt"

type Person struct {
	LastName  string
	FirstName string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s: 年齢%d歳", p.LastName, p.FirstName, p.Age)
}

func main() {
	keisuke := Person{
		LastName:  "Keisuke",
		FirstName: "Honda",
		Age:       30,
	}
	fmt.Println(keisuke.String())
}
