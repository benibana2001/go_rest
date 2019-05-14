package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/benibana2001/go_rest/data"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

type Controller struct{}

// users
func (c Controller) GetUsers (w http.ResponseWriter, r *http.Request) {
	users := selectAllUser()
	fmt.Fprintf(w, "%+v", users)
}

// create user
func (c Controller) CreateUser (w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	createUser(r)
}

// user
func (c Controller) GetUser (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	user := selectUser(id)
	fmt.Fprintf(w, "%+v", user)
}

// update user
func (c Controller) UpdateUser (w http.ResponseWriter, r *http.Request) {
	// not implemented yet
}

// Create user
func createUser(r *http.Request) {
	db := data.ConnectDb()
	defer db.Close()

	user := parseJson(r)
	db.Create(&user)
	fmt.Printf("newUser is %v\n", user)
}

// Select user
func selectUser(id string) data.User {
	db := data.ConnectDb()
	defer db.Close()

	var user data.User
	db.First(&user, "id = ?", id)

	// todo: Marshal to JSON
	return user
}

// Select all user
func selectAllUser() []data.User {
	db := data.ConnectDb()
	defer db.Close()

	// Select
	var users []data.User
	db.Find(&users)

	// todo: Marshal to JSON
	return users
}

// Parse Json
func parseJson(r *http.Request) data.User {
	// Parse
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(5)
	}
	fmt.Println("Read : ", body)

	// Unmarshal
	user := data.User{}
	err = json.Unmarshal(body, &user)
	fmt.Println("Unmarshal : ", &user)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(4)
	}
	return user
}
