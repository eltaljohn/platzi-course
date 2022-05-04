package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("ftEmployee: %v\n", ftEmployee)
	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
