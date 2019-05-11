package main

import (
	"github.com/benibana2001/go_rest/data"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func dropUserTable(db *gorm.DB) {
	db.DropTable(&data.User{})
}

func main() {
	db := data.ConnectDb()
	dropUserTable(db)
}
