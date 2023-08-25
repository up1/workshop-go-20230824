package demo_test

import (
	"demo"
	"demo/db"
	"testing"
)

func TestProcessWithSuccess(t *testing.T) {
	ur := db.NewUserRepository(nil)
	d2 := demo.NewDemoBusiness(ur)
	r := d2.Process(1)
	if !r {
		t.Error("Test fail with success")
	}
}
