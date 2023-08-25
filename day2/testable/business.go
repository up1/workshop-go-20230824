package demo

type DemoBusiness struct {
}

func (d DemoBusiness) Process() {
	println("TODO Next")
}

func NewDemoBusiness() *DemoBusiness {
	return &DemoBusiness{}
}
