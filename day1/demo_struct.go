package main

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
func (e *Employee) leave() bool {
	e.Name = "Update somkiat"
	return true
}

func main() {
	e := EmployeeWithEmail{}
	e.leave()
	e.Email = "email"
	println(e.Name)

	var d DayOfMonth = 5 // Day of month
	println(d)
}
