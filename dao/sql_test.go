package dao

import (
	"fmt"
	"testing"
)

func TestQueryAuthor(t *testing.T) {
	author, err := QueryAuthor("222@qq.com", "2222")
	if err != nil {
		t.Error(err)
	}
	t.Log(author)
	fmt.Println(author)
}
