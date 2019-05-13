package main

import (
	"github.com/benibana2001/go_rest/Controllers"
	"github.com/benibana2001/go_rest/data"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	db := data.ConnectDb()
	defer db.Close()

	// Listen
	controller := Controllers.Controller{}
	http.HandleFunc("/users", controller.HUsers)

	http.HandleFunc("/users/", controller.HUser)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
