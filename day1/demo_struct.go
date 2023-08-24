package main

type DayOfMonth int

type Employee struct {
	Id   string
	Name string
	Age  int
}

// Receiver
// 1. Value Receiver
// 2. Point Receiver
func (e *Employee) leave() bool {
	e.Name = "Update somkiat"
	return true
}

func main() {
	e := Employee{Name: "somkiat"}
	e.leave()
	println(e.Name)

	var d DayOfMonth = 5 // Day of month
	println(d)
}
