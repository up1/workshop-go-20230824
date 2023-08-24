package main

type reporter interface {
	process() bool
}

type reportA struct{}

func (r reportA) process() bool {
	println("Called report A")
	return true
}

func executeReport(r reporter) {
	r.process()
}

func main() {
	ra := reportA{}
	executeReport(ra)
}
