package main

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	controller.UseMidWare(router)
	controller.UseSessions(router)

	controller.SetStatics(router)
	controller.LoadHTML(router)

	controller.LoadPages(router)
	controller.LoadControllers(router)

	router.Run()
}
