package main

import (
	"github.com/benibana2001/go_rest/Controllers"
	"github.com/benibana2001/go_rest/data"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"time"
)

func main() {
	db := data.ConnectDb()
	defer db.Close()

	// Listen
	controller := Controllers.Controller{}
	router := mux.NewRouter()

	router.HandleFunc("/users/", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users/", controller.CreateUser).Methods("POST")

	router.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controller.UpdateUser).Methods("POST")

	srv := &http.Server {
		Handler: router,
		Addr:    "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatalln(srv.ListenAndServe())
}
