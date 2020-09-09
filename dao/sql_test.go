package dao

import (
	"fmt"
	"strconv"
	"testing"
)

func TestQueryAuthor(t *testing.T) {
	author, err := QueryAuthorByEmailAndPasswd("222@qq.com", "2222")
	if err != nil {
		t.Error(err)
	}
	t.Log(author)
	fmt.Println(author)
}
func TestQueryAuthorById(t *testing.T) {
	author, err := QueryAuthorByID(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(author)
	fmt.Println(author)
	fmt.Println(AddAuthor(author))
}
func TestQueryMapIDByArticleIDAndAuthorID(t *testing.T) {
	id := QueryMapIDByArticleIDAndAuthorID(6, 2)
	fmt.Println(id)

	a := "1"
	b, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(b)
}