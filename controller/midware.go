package controller

import (
	"blog/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"fmt"
)

// UseMidWare 配置路由中间件
func UseMidWare(router *gin.Engine) {
	router.Use(PasswordMidWare())
}
// PasswordMidWare 密码加密中间件
func PasswordMidWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.PostForm("password")) != 0 {
			enc := service.EncodePassword(c.PostForm("password"))
			c.Set("password", enc)
			fmt.Println(enc)
		}
		if len(c.PostForm("repassword")) != 0 {
			enc := service.EncodePassword(c.PostForm("password"))
			c.Set("repassword", enc)
			fmt.Println(enc)
		}
	}
}

// UseSessions 使用session初始化
func UseSessions(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("loginsession", store))
}
