package demo

import (
	"demo/db"

	"github.com/jackc/pgx/v5"
)

type DemoBusiness struct {
	conn *pgx.Conn
}

func (d DemoBusiness) Process(id int) {
	println("TODO Next")
	db.GetUserById(d.conn, id)
}

func NewDemoBusiness(conn *pgx.Conn) *DemoBusiness {
	return &DemoBusiness{conn: conn}
}
