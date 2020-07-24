package controller

import (
	"blog/model"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"encoding/json"
)

// SetStatics 设置静态文件目录
func SetStatics(router *gin.Engine) {
	router.Static("/css", "./statics/css")
	router.Static("/img", "./statics/img")
	router.Static("/scss", "./statics/scss")
	router.Static("/vendor", "./statics/vendor")
	router.Static("/js", "./statics/js")
	router.Static("/fonts", "./statics/fonts")
	router.Static("/images", "./statics/images")
	router.Static("/languages", "./statics/languages")
	router.Static("/lib", "./statics/lib")
	router.Static("/plugins", "./statics/plugins")
	router.Static("/src", "./statics/src")
	router.Static("/picture", "./picture")
	router.StaticFile("test.md", "./statics/test.md")
	router.StaticFile("editormd.js", "./statics/js/editormd.js")
	router.StaticFile("favicon.ico", "./statics/img/favicon.ico")
}

// LoadHTML 加载html模板文件
func LoadHTML(router *gin.Engine) {
	router.LoadHTMLGlob("./view/*")
}

// LoadPages 映射路径和页面
func LoadPages(router *gin.Engine) {
	router.GET("/", IndexPage)
	router.GET("/index", IndexPage)
	router.GET("/login", LoginPage)
	router.GET("/home", HomePage)
	router.GET("/registe", RegistePage)
	router.GET("/lostpasswd", LostPasswdPage)
	router.GET("/editor", EditorPage)
	router.GET("/error", ErrPage)
	router.GET("/about", AboutPage)
}

// IndexPage 返回idex页面
func IndexPage(c *gin.Context) {
	fmt.Println("index:" + c.Request.URL.RequestURI())
	// c.Header("CacheControl", "no-cache")
	session := sessions.Default(c)
	var author model.Author
	if session.Get("author") != nil {
		// author = session.Get("author").(model.Author)
		// fmt.Printf("%T", session.Get("author"))
		json.Unmarshal([]byte(session.Get("author").(string)), &author)
	}
	// fmt.Println(author)
	var home = "/home"
	var nick = ""
	if len(author.Email) != 0 {
		home = fmt.Sprintf("%s?id=%d", home, author.ID)
		nick = author.Nick
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"nick": nick,
		"login":  "/login",
		"home":   home,
		"registe": "/registe",
		"logout": "/logout",
	})

}

// LoginPage 返回登陆界面
func LoginPage(c *gin.Context) {
	var email, _ = c.GetPostForm("email")
	// fmt.Println(email)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"email": email,
		"login":   "/logincontroller",
		"lost":  "/lostpasswd",
	})
}

// HomePage 返回用户主界面
func HomePage(c *gin.Context) {

	var author model.Author
	session := sessions.Default(c)
	if session.Get("author") != nil {
		json.Unmarshal([]byte(session.Get("author").(string)), &author)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/error?code=403")
	}

	c.HTML(200, "home.html", gin.H{
		"author": author,
		"editor": "/editor",
		"about": "/about",
		"headpicture": "picture/2a679803-06c5-4599-8d57-0e5387a8dfc1.JPG",
	})	
}

// LostPasswdPage 返回忘记密码的页面和用户登陆注册页面相同
func LostPasswdPage(c *gin.Context) {
	c.HTML(http.StatusOK, "losspasswd.html", gin.H{
		"url": "/lostcontroller",
		"message": "input a new password",
	})
}
// RegistePage 注册页面
func RegistePage(c *gin.Context) {
	value, ok := c.Get("email")
	var email = ""
	if ok {
		email = value.(string)
	}
	c.HTML(http.StatusOK, "registe.html", gin.H{
		"url": "/registecontroller",
		"email": email,

	})
}
// ErrPage 出错页面
func ErrPage(c *gin.Context) {
	code := c.Query("code")
	c.HTML(http.StatusOK, "error.html", gin.H{
		"code": code,
		"home": "/home",
		"login": "/login",
		"registe": "/registe",
	})
}

// EditorPage 返回编辑器页面，使用markdown
func EditorPage(c *gin.Context) {
	// code := c.Query("code")
	c.HTML(http.StatusOK, "editor.html", nil)
}
// AboutPage 关于作者信息
func AboutPage(c *gin.Context) {

}