package main

import "fmt"

type DayOfMonth int

type Employee struct {
	Id   string
	Name string
	Age  int
}
type EmployeeWithEmail struct {
	Employee
	Email string
}

// Receiver
// 1. Value Receiver
// 2. Pointer Receiver
func (e Employee) leave() bool {
	e.Name = "Update somkiat"
	return true
}

func (e Employee) String() string {
	return "........."
}

func main() {
	e := Employee{}
	e.leave()
	fmt.Println(e)

	var d DayOfMonth = 5 // Day of month
	println(d)
}
