package main

import (
	"bytes"
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
		url: "http://localhost:8081/users/",
		contentType: "application/json",
	}
	//user := data.User{Id: 0, Name: "Jeff Bezos", Email: "bezos@mail.com"}
	//buf, _ := json.Marshal(&user)
	//c.post(buf)
	c.get()
}

func (c Client) get () {
	resp, err := http.Get(c.url)
	if err != nil {
		fmt.Printf("Error : %v", err)
		os.Exit(2)
	}
	fmt.Println(resp)
}

func (c Client) post (bs []byte) {
	fmt.Println("[]byte is -> ", bs)
	_, err := http.Post(c.url, c.contentType, bytes.NewBuffer(bs))
	fmt.Printf("body (io.Reader) is -> %v \n", bytes.NewBuffer(bs))

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(3)
	}

}

func (c Client) delete () {
	//resp, err := http.Get(c.url)
	//if err != nil {
	//	fmt.Printf("Error : %v", err)
	//	os.Exit(2)
	//}

}