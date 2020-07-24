package dao

import (
	"fmt"
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
