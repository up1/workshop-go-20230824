package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer conn.Close(context.Background())
	return conn
}

type User struct {
	Id    int
	Email string
}

func GetUserById(conn *pgx.Conn, id int) (User, error) {
	err := conn.QueryRow(context.Background(), "select * from users where id=$1").Scan(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return User{}, nil
}
