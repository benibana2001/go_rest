package main

import (
	"github.com/benibana2001/go_rest/data"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func createUserTable(db *gorm.DB) {
	db.CreateTable(&data.User{})
	// Seeding -Create columns
	u1 := data.User{Id: 0, Name: "Takeru Satou", Email: "takeru@mail.jp"}
	u2 := data.User{Id: 0, Name: "Hanako Yamada", Email: "hanako@mail.jp"}
	u3 := data.User{Id: 0, Name: "Satoshi Tajima", Email: "satoshi@mail.jp"}
	db.Create(&u1)
	db.Create(&u2)
	db.Create(&u3)
}

func main()  {
	db := data.ConnectDb()
	createUserTable(db)
}
