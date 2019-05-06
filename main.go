package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"strings"
)

type User struct {
	//gorm.Model
	Id int
	Name string
	Email string
}

func main() {
	db, err := connectDb()
	defer db.Close()
	if err != nil{
		fmt.Printf("%v", err)
		panic("failed to connect database")
	}

	// Read
	var user User
	var users []User
	//db.First(&user, "id = ?", "1") // find product with code l1212


	// Handling
	hUsers := func(w http.ResponseWriter, r *http.Request) {
		users = selectAllUser(db)
		fmt.Printf("Current URL is %v\n", r.URL)
		fmt.Fprintf(w, "%+v", users)
	}
	hUser := func(w http.ResponseWriter, r *http.Request) {
		s := strings.Split(r.URL.Path, "/")// ["" "users" "1"]
		user = selectUser(db, s[2])

		// Todo: request_id > len(data)

		fmt.Printf("Current Path is %q Type is %T\n", s, s)
		fmt.Printf("userId is %v", s[2])
		fmt.Fprintf(w, "%+v", user)
	}

	// Listen
	http.HandleFunc("/users", hUsers)
	http.HandleFunc("/users/", hUser)
	fmt.Printf("handler is %T", hUser)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Connect Database
func connectDb() (*gorm.DB, error){
	fmt.Println("Connecting database...")
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3308)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	//should not defer
	return db, nil
}

// Select user
func selectUser(db *gorm.DB, id string) User{
	var user User
	db.First(&user, "id = ?", id)
	return user
}

// Select all user
func selectAllUser(db *gorm.DB) []User{
	// Select
	var users []User
	db.Find(&users)
	return users
}

