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
	_, err := d.ur.GetUserById(id)
	if err != nil {
		return false
	}
	return true
}

func NewDemoBusiness(ur db.MyRepository) *DemoBusiness {
	return &DemoBusiness{ur: ur}
}
