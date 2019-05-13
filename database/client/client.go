package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/benibana2001/go_rest/data"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"os"
)

type Client struct {
	url string
	contentType string
	db *gorm.DB
}

func main() {
	db := data.ConnectDb()
	defer db.Close()

	c := Client{
		url: "http://localhost:8081/users",
		contentType: "application/json",
	}
	user := data.User{Id: 0, Name: "Jeff Bezos", Email: "bezos@mail.com"}
	buf, _ := json.Marshal(&user)
	c.post(buf)
}

func (c Client) get () {
	resp, err := http.Get(c.url)
	if err != nil {
		fmt.Printf("Error : %v", err)
		os.Exit(2)
	}
	fmt.Println(resp)
}

func (c Client) post (json []byte) {
	fmt.Println("[]byte is -> ", json)
	_, err := http.Post(c.url, c.contentType, bytes.NewBuffer(json))
	fmt.Printf("[]byte is -> %v", bytes.NewBuffer(json))

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(3)
	}

}
