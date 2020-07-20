package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.LoadHTMLFiles("./statics/index.html")
	router.LoadHTMLGlob("./view/*")
	router.Static("/css", "./statics/css")
	router.Static("/img", "./statics/img")
	router.Static("/scss", "./statics/scss")
	router.Static("/vendor", "./statics/vendor")
	router.Static("/js", "./statics/js")
	router.Static("/fonts", "./statics/fonts")
	router.Static("/images", "./statics/images")
	router.GET("/", indexPage)
	router.GET("/login", loginPage)
	router.POST("/login", loginPage)
	router.POST("/home", homePage)
	router.GET("/lost", lostPage)
	router.Run()
}
func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "/login")
}
func loginPage(c *gin.Context) {
	var email, _ = c.GetPostForm("email")
	fmt.Println(email)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"email": email,
		"url":   "/home",
		"lost":  "/lost",
	})
}
func homePage(c *gin.Context) {
	// var email, _ = c.GetPostForm("email")
	// var passwd, _ = c.GetPostForm("password")
	c.HTML(200, "home.html", gin.H{
		"name": "宝贝",
		"home": "/home",
	})
}
func lostPage(c *gin.Context) {

}
