package main

import (
	//"encoding/json"
	"fmt"
	"github.com/benibana2001/go_rest/data"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Parse Json

func parseJson(r *http.Request) data.User{

	i, _ := strconv.Atoi(r.FormValue("Id"))
	n := r.FormValue("Name")
	e := r.FormValue("Email")
	newUser := data.User{
		Id:    i,
		Name:  n,
		Email: e,
	}
	return newUser
}

func main() {
	db := data.ConnectDb()
	defer db.Close()

	// Handling
	hUsers := func(w http.ResponseWriter, r *http.Request) {
		m := r.Method
		if m == "GET" {
			users := selectAllUser(db)
			fmt.Fprintf(w, "%+v", users)
		} else if m == "POST" {
			newUser := parseJson(r)
			fmt.Printf("newUser is %v\n", newUser)
			db.Create(&newUser)
		}
	}
	hUser := func(w http.ResponseWriter, r *http.Request) {
		s := strings.Split(r.URL.Path, "/")// ["" "users" "1"]

		// Todo: request_id > len(data)

		// Request method is GET -> selectUser, PUT -> updateUser, Delete -> deleteUser

		m := r.Method
		if m == "GET" {
			// Select user
			user := selectUser(db, s[2])
			fmt.Fprintf(w, "%+v", user)
		} else if m == "POST" {
			// Todo: Implement Update user
			//user := selectUser(db, s[2])
		}
	}

	// Listen
	http.HandleFunc("/users", hUsers)

	http.HandleFunc("/users/", hUser)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Select user
func selectUser(db *gorm.DB, id string) data.User{
	var user data.User
	db.First(&user, "id = ?", id)

	// todo: Marshal to JSON
	return user
}

// Select all user
func selectAllUser(db *gorm.DB) []data.User{
	// Select
	var users []data.User
	db.Find(&users)

	// todo: Marshal to JSON
	return users
}
