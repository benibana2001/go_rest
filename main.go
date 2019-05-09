package main

import (
	//"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"strconv"

	//"strconv"
	"strings"
)

type User struct {
	//gorm.Model
	Id    int `gorm:"primary_key"`
	Name  string
	Email string
}

// Sample table to test migration
type Sample struct {
	Id int
	Column1 string
	Column2 int
}

func createSampleTable(db *gorm.DB) {
	db.CreateTable(&Sample{})

	// Seeding - Create columns
	s1 := Sample{Id: 0, Column1: "Hello!", Column2: 100}
	s2 := Sample{Id: 0, Column1: "Good Morning!", Column2: 50}
	db.Create(&s1)
	db.Create(&s2)
	samples := selectAllSample(db)
	fmt.Printf("%+v", samples)
}

func createUserTable(db *gorm.DB) {
	db.CreateTable(&User{})
	// Seeding -Create columns
	u1 := User{Id: 0, Name: "Takeru Satou", Email: "takeru@mail.jp"}
	u2 := User{Id: 0, Name: "Hanako Yamada", Email: "hanako@mail.jp"}
	u3 := User{Id: 0, Name: "Satoshi Tajima", Email: "satoshi@mail.jp"}
	db.Create(&u1)
	db.Create(&u2)
	db.Create(&u3)
}

func dropUserTable(db *gorm.DB) {
	db.DropTable(&User{})
}

// Parse Json

func parseJson(r *http.Request) User{

	i, _ := strconv.Atoi(r.FormValue("Id"))
	n := r.FormValue("Name")
	e := r.FormValue("Email")
	newUser := User{
		Id:    i,
		Name:  n,
		Email: e,
	}
	return newUser
}

func main() {
	db, err := connectDb()
	defer db.Close()
	if err != nil{
		fmt.Printf("%v", err)
		panic("failed to connect database")
	}

	var user User
	var users []User

	// Handling
	hUsers := func(w http.ResponseWriter, r *http.Request) {
		m := r.Method
		if m == "GET" {
			users = selectAllUser(db)
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
			user = selectUser(db, s[2])
			fmt.Fprintf(w, "%+v", user)
		} else if m == "POST" {
			// Update user
			user = selectUser(db, s[2])
		}
	}

	// Listen
	http.HandleFunc("/users", hUsers)

	http.HandleFunc("/users/", hUser)

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

// Select all sample
func selectAllSample(db *gorm.DB) []Sample{
	// Select
	var samples []Sample
	db.Find(&samples)
	return samples
}

func createUser() {}

func updateUser() {}

func deleteUser() {}
