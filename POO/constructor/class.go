package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

// Constructor
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// 1
	e := Employee{}
	fmt.Printf("%v\n", e)
	// 2
	e2 := Employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)
	// 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 2
	e3.name = "Name 3"
	fmt.Printf("%v\n", *e3)
	//4
	e4 := NewEmployee(1, "John", true)
	fmt.Printf("%v\n", *e4)

}
