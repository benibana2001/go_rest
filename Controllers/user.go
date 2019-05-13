package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/benibana2001/go_rest/data"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Controller struct{}

// Handler functions
func (c Controller) HUsers (w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == "GET" {
		users := selectAllUser()
		fmt.Fprintf(w, "%+v", users)
	} else if m == "POST" {
		defer r.Body.Close()
		createUser(r)
	}
}
func (c Controller) HUser (w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/") // ["" "users" "1"]

	// Todo: request_id > len(data)

	// Request method is GET -> selectUser, PUT -> updateUser, Delete -> deleteUser

	m := r.Method
	if m == "GET" {
		// Select user
		user := selectUser(s[2])
		fmt.Fprintf(w, "%+v", user)
	} else if m == "POST" {
		// Todo: Implement Update user
		//user := selectUser(db, s[2])
	}
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
