package demo_test

import (
	"demo"
	"demo/db"
	"fmt"
	"testing"
)

type mockUserRepository struct {
	status bool
}

func (m *mockUserRepository) GetUserById(id int) (db.User, error) {
	if m.status {
		return db.User{}, nil
	}
	return db.User{}, fmt.Errorf("Database error")
}

func TestProcessWithSuccess(t *testing.T) {
	ur := mockUserRepository{true}
	d2 := demo.NewDemoBusiness(&ur)
	r := d2.Process(1)
	if !r {
		t.Error("Test fail with success")
	}
}

func TestProcessWithError(t *testing.T) {
	ur := mockUserRepository{false}
	d2 := demo.NewDemoBusiness(&ur)
	r := d2.Process(1)
	if r {
		t.Error("Test fail with success")
	}
}
