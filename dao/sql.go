package dao

import (
	"blog/model"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	user     = "worker"
	passwd   = "Hh123456"
	ip       = "47.92.139.92"
	port     = "3306"
	database = "test"
)
var (
	db *sql.DB
)

func init() {
	var url string
	url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, passwd, ip, port, database)
	var err error
	db, err = sql.Open("mysql", url)
	if err != nil {
		log.Panic(err)
	}
}

/**
*  通过email和passwd获取mysql中的作者信息
**/
func QueryAuthor(email string, passwd string) (author model.Author, err error) {
	stmt, err := db.Prepare("SELECT * FROM author WHERE EMAIL = ? && PASSWD = ?")
	if err != nil {
		log.Println(err)
		return
	}
	row := stmt.QueryRow(email, passwd)
	row.Scan(&author.ID, &author.Nick, &author.Passwd, &author.Email, &author.Description, &author.Picture)
	return
}
func CloseConnection() {
	if db == nil {
		return
	}
	db.Close()
}
