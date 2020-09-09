package controller

import (
	"blog/model"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"encoding/json"
	"blog/service"
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
	router.Static("/article", "./article")
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
	router.GET("/getarticle", ArticlePage)
}
// ArticlePage 显示文章页
func ArticlePage(c *gin.Context) {
	id := c.Query("articleid")
	if len(id) == 0 {
		c.Redirect(http.StatusMovedPermanently, "/error?code=404")
	}
	articleid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusMovedPermanently, "/error?code=304")
	}
	log.Println("ArticlePage", "articleid:" , articleid)
	author := service.GetAuthorBySession(c)
	log.Println("ArticlePage", "author:", author)
	if !service.VertifyArticleByArticleIDAndAuthorID(author.ID, articleid) {
		c.Redirect(http.StatusMovedPermanently, "/error?code=304")
	}
	article, err := service.GetArticleByID(articleid)
	if err != nil {
		log.Println("ArticlePage", err)
		c.Redirect(http.StatusMovedPermanently, "/error?code=404")
	}
	log.Println("ArticlePage", "article:", article)
	c.HTML(http.StatusOK, "showarticle.html", gin.H{
		"author":  author,
		"article": article,
		"modify":  "/modify",
		"home":    "/home",
	})
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
	// var nick = ""
	// if len(author.Email) != 0 {
	// 	home = fmt.Sprintf("%s?id=%d", home, author.ID)
	// 	nick = author.Nick
	// }
	c.HTML(http.StatusOK, "index.html", gin.H{
		"nick": author.Nick,
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
	articleids := service.GetArticleIDListByAuthorID(author.ID)
	fmt.Println(articleids)

	arraynum := len(articleids)/3
	if len(articleids)%3 != 0 {
		arraynum ++
	}

	articlesmap := make([][]model.Article, arraynum)
	var idx = 0
	for i := 0; i < arraynum; i ++ {
		articlesmap[i] = make([]model.Article, 0)
		for j := 0; j < 3 && idx < len(articleids); j ++ {
			article, err := service.GetArticleByID(articleids[idx])
			idx ++
			if err != nil {
				continue
			}
			articlesmap[i] = append(articlesmap[i], article)
		}
	}

	c.HTML(200, "home.html", gin.H{
		"author": author,
		"editor": "/editor",
		"about": "/about",
		"headpicture": author.Picture,
		"articlesmap": articlesmap,
		"articleurl": "/getarticle",
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
	var author model.Author
	session := sessions.Default(c)
	if session.Get("author") != nil {
		json.Unmarshal([]byte(session.Get("author").(string)), &author)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/error?code=403")
	}
	c.HTML(http.StatusOK, "editor.html", gin.H{
		"url": "/savearticle",
		"nick": author.Nick,
		"home": "/home",
	})
}
// AboutPage 关于作者信息
func AboutPage(c *gin.Context) {

}