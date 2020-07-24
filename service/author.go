package service

import (
	"blog/model"
	"blog/dao"
	"log"
)
// GetAuthorByEmailAndPasswd 通过email和passwd获取对应author
func GetAuthorByEmailAndPasswd(email string, passwd string) model.Author {
	author, err := dao.QueryAuthorByEmailAndPasswd(email, passwd)
	if err != nil {
		log.Println(err)
	}
	return author
}
// GetAuthorByID 通过ID获取作者
func GetAuthorByID(id int64) model.Author {
	author, err := dao.QueryAuthorByID(id)
	if err != nil {
		log.Println(err)
	}
	return author
}
// RegisteAuthor 注册用户
func RegisteAuthor(author model.Author) int64 {
	return dao.AddAuthor(author)
}
// VertifyAuthorByNicknameAndEmail 验证用户是否存在
func VertifyAuthorByNicknameAndEmail(nick, email string) bool {
	author, err := dao.QueryAuthorByNickAndEmail(nick, email)
	if err != nil {
		return false
	}
	if author.Email != email {
		return false
	}
	return true
}
// UpdatePasswordByEmail 更新密码
func UpdatePasswordByEmail(email, password string) int64 {
	return dao.UpdateAuthorPasswordForEmail(email, password)
}