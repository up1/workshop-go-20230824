package demo_test

import (
	"demo"
	"demo/db"
	"testing"
)

type mockUserRepository struct {
}

func (m *mockUserRepository) GetUserById(id int) (db.User, error) {
	return db.User{}, nil
}

func TestProcessWithSuccess(t *testing.T) {
	ur := mockUserRepository{}
	d2 := demo.NewDemoBusiness(&ur)
	r := d2.Process(1)
	if !r {
		t.Error("Test fail with success")
	}
}
