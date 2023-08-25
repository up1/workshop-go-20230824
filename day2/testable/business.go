package demo

import (
	"demo/db"

	"github.com/jackc/pgx/v5"
)

var Connection *pgx.Conn

type DemoBusiness struct {
	ur *db.UserRepository
}

func (d DemoBusiness) Process(id int) {
	println("TODO Next")
	d.ur.GetUserById(id)
}

func NewDemoBusiness(ur *db.UserRepository) *DemoBusiness {
	return &DemoBusiness{ur: ur}
}
