package demo

import (
	"demo/db"

	"github.com/jackc/pgx/v5"
)

var Connection *pgx.Conn

type DemoBusiness struct {
	ur db.MyRepository
}

func (d DemoBusiness) Process(id int) bool {
	println("TODO Next")
	d.ur.GetUserById(id)
	return true
}

func NewDemoBusiness(ur db.MyRepository) *DemoBusiness {
	return &DemoBusiness{ur: ur}
}
