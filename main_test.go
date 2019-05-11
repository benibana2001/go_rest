package main

import (
	"github.com/benibana2001/go_rest/data"
	"testing"
)

func TestSelectUser(t *testing.T) {
	db := data.ConnectDb()
	defer db.Close()

	user := selectUser(db, "1")
	if user.Id != 1 {
		t.Errorf("Expected user.Id is 1, but got %v", user.Id)
	}
	if user.Name != "Takeru Satou" {
		t.Errorf("Expected user.Name is 'Takeru Satou', but got '%v'", user.Name)
	}
	if user.Email != "takeru@mail.jp" {
		t.Errorf("Expected user.Email is 'takeru@mail.jp', but got '%v'", user.Email)
	}
}

func TestSelectAllUser(t *testing.T) {
	db := data.ConnectDb()
	defer db.Close()

	users := selectAllUser(db)
	if users[0].Id != 1 {
		t.Errorf("Expected users[0].Id is 1, but got %v", users[0].Id)
	}
	if users[2].Id != 3 {
		t.Errorf("Expected users[2].Id is 3, but got %v", users[2].Id)
	}
}
