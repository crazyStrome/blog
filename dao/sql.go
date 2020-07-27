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

// QueryAuthorByEmailAndPasswd 通过email和passwd获取mysql中的作者信息
func QueryAuthorByEmailAndPasswd(email string, passwd string) (author model.Author, err error) {
	stmt, err := db.Prepare("SELECT * FROM author WHERE EMAIL = ? && PASSWD = ?")
	if err != nil {
		log.Println(err)
		return
	}
	row := stmt.QueryRow(email, passwd)
	row.Scan(&author.ID, &author.Nick, &author.Passwd, &author.Email, &author.Description, &author.Picture)
	return
}
// QueryAuthorByNickAndEmail 通过昵称和邮箱获取作者
func QueryAuthorByNickAndEmail(nick string, email string) (author model.Author, err error) {
	stmt, err := db.Prepare("SELECT * FROM author WHERE nickname = ? && email = ?")
	if err != nil {
		log.Println(err)
		return
	}
	row := stmt.QueryRow(nick, email)
	row.Scan(&author.ID, &author.Nick, &author.Passwd, &author.Email, &author.Description, &author.Picture)
	return
}
// UpdateAuthorPasswordForEmail 通过邮箱更新信息
func UpdateAuthorPasswordForEmail(email, password string) int64 {
	stmt, err := db.Prepare("update author set passwd = ? where email = ?")
	if err != nil {
		log.Println(err)
		return -1
	}
	result, err := stmt.Exec(password, email)
	if err != nil {
		return -1
	}
	id, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return id
}

// QueryAuthorByID 通过作者id获取作者信息
func QueryAuthorByID(id int64) (author model.Author, err error) {
	stmt, err := db.Prepare("SELECT * FROM author WHERE id = ?")
	if err != nil {
		log.Println(err)
		return
	}
	row := stmt.QueryRow(id)
	row.Scan(&author.ID, &author.Nick, &author.Passwd, &author.Email, &author.Description, &author.Picture)
	return
}

// AddAuthor 添加作者信息
func AddAuthor(author model.Author) int64 {
	stmt, err := db.Prepare("INSERT INTO author(nickname, passwd, email, description, picture) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return -1
	}
	result, err := stmt.Exec(author.Nick, author.Passwd, author.Email, author.Description, author.Picture)
	if err != nil {
		return -1
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return id
}

// QueryArticleByID 通过文章id获取文章具体信息
func QueryArticleByID(id int64) (article model.Article, err error) {
	stmt, err := db.Prepare("SELECT * FROM article WHERE id = ?")
	if err != nil {
		log.Println(err)
		return
	}
	row := stmt.QueryRow(id)
	row.Scan(&article.ID, &article.Title, &article.Destination, &article.Timestep)
	return
}

// QueryArticleByTitle 通过文章名查询文章信息
func QueryArticleByTitle(title string) (article model.Article, err error) {
	stmt, err := db.Prepare("SELECT * FROM article WHERE title = ?")
	if err != nil {
		log.Println(err)
		return
	}
	row := stmt.QueryRow(title)
	row.Scan(&article.ID, &article.Title, &article.Destination, &article.Timestep)
	return
}

// AddArticle 添加文章到数据库
func AddArticle(article model.Article) int64 {
	stmt, err := db.Prepare("INSERT INTO article(title, destination, timestep) values(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return -1
	}
	result, err := stmt.Exec(article.Title, article.Destination, article.Timestep)
	if err != nil {
		return -1
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return id
}

// QueryArticleIDsByAuthorID 通过作者id获取文章列表
func QueryArticleIDsByAuthorID(authorID int64) []int64 {
	stmt, err := db.Prepare("SELECT articleid FROM map_author_article WHERE authorid = ?")
	if err != nil {
		log.Println(err)
		return []int64{}
	}
	var res = []int64{}
	rows, err := stmt.Query(authorID)
	for rows.Next() {
		var id int64
		rows.Scan(&id)
		res = append(res, id)
	}
	return res
}

// AddMapAuthorIDAndArticleID 添加作者和文章的映射
func AddMapAuthorIDAndArticleID(authorID, articleID int64) int64 {
	stmt, err := db.Prepare("INSERT INTO map_author_article(authorid, articleid) values(?, ?)")
	if err != nil {
		log.Println(err)
		return -1
	}
	result, err := stmt.Exec(authorID, articleID)
	if err != nil {
		return -1
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return id
}

// CloseConnection 关闭db
func CloseConnection() {
	if db == nil {
		return
	}
	db.Close()
}
