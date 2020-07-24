package model

import "time"
// Author 作者具体信息
type Author struct {
	ID          int `json:"id"`
	Nick        string `json:"nick"`
	Passwd      string `json:"password"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}
// Article 文章具体信息
type Article struct {
	ID          int 
	Title       string
	Destination string
	Timestep    time.Time
}
