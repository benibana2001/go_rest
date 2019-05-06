package main

import (
	"fmt"
	"testing"
)

func TestSelectUser(t *testing.T) {
	db, err := connectDb()
	defer db.Close()
	if err != nil{
		fmt.Printf("%v", err)
	}

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
