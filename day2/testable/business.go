package demo

import xxx "demo/db"

type DemoBusiness struct {
}

func (d DemoBusiness) Process() {
	xxx.XXX()
	println("TODO Next")
}

func NewDemoBusiness() *DemoBusiness {
	return &DemoBusiness{}
}
