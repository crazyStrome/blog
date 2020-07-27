package model

import "time"
// Author 作者具体信息
type Author struct {
	ID          int64 `json:"id"`
	Nick        string `json:"nick"`
	Passwd      string `json:"password"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}
// Article 文章具体信息
type Article struct {
	ID          int64 
	Title       string
	Destination string
	Timestep    time.Time
}
// ResponseDAO 和 DataDAO 用于解析SMMS返回的json
type DataDAO struct {
	Filename string `json: "filename"`
	URL string `json: "url"`
}
type ResponseDAO struct {
	Success bool `json: "success"`
	Code string `json: "code"`
	Message string `json: "message"`
	Data DataDAO `json: "data"`
}