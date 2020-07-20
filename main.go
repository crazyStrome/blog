package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(logger())
	router.LoadHTMLFiles("./statics/index.html")
	router.StaticFile("/css", "./css")
	router.StaticFile("/img", "./img")
	router.StaticFile("/scss", "./scss")
	router.StaticFile("/vendor", "./vendor")
	router.GET("/", func(c *gin.Context) {
		// c.HTML(http.StatusOK, "index.html", "sss")
		c.String(200, "宝贝😍，嘻嘻~")
	})
	router.Run()
}
func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("1111")
		c.Next()
		fmt.Println("2222")
	}
}
