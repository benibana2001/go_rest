package main

import (
	"encoding/json"
	"fmt"
	"github.com/benibana2001/go_rest/data"
	"io"
	"net/http"
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

func TestParesJson() {
	user := &data.User{
		Id: 0,
		Name: "Jeff Bezos",
		Email: "bezos@mail.com",
	}
	json, _ := json.Marshal(&user)
	bs := []byte(json)
	postClient, _ := http.Post(
		"http://localhost:8081/users",
		"application/json",
		logReader{},
		)
	postClient.post

	parseJson

}

type logReader struct {}

func (lr logReader) Read(p []byte) (n int, err error) {
	fmt.Println(n)
	return n, err
}
