package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

type User struct {
	//gorm.Model
	Id    int    `json:"Id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

// Connect Database
func ConnectDb() (*gorm.DB){
	fmt.Println("Connecting database...")
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3308)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Error was occur : ", err)
		os.Exit(1)
	}
	//should not defer
	return db
}
