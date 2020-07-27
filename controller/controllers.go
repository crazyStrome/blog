package controller

import (
	"time"
	// "io/ioutil"
	// "io"
	"blog/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"blog/service"
	"net/http"
	"fmt"
	"encoding/json"
	"strings"
	// "os"
)
// LoadControllers 加载所有controller
func LoadControllers(router *gin.Engine) {
	router.GET("/logout", LogoutController)

	router.POST("/logincontroller", LoginController)
	router.POST("/registecontroller", RegisteController)
	router.POST("/lostcontroller", LostPasswordController)
	router.POST("/upload", UploadController)
	router.POST("/savearticle", SaveArticle)
}
// SaveArticle 储存文章
func SaveArticle(c *gin.Context) {
	content := c.PostForm("test-editormd-markdown-doc")
	title := c.PostForm("title")

	var author model.Author
	session := sessions.Default(c)
	if session.Get("author") != nil {
		json.Unmarshal([]byte(session.Get("author").(string)), &author)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/error?code=403")
		return
	}
	filename, err := service.SaveArticle(content)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/error?code=503")
		return
	}
	article := model.Article{
		Title: title,
		Destination: filename,
		Timestep: time.Now(),
	}
	articleid := service.AddArticle(article)
	if articleid < 0 {
		c.Redirect(http.StatusMovedPermanently, "/error?code=403")
		return
	}
	line := service.MapAuthorIDAndArticleID(author.ID, articleid)
	if line < 0 {
		c.Redirect(http.StatusMovedPermanently, "/error?code=403")
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/home")
}
// UploadController markdown上传图片处理器
func UploadController(c *gin.Context) {
	form, _ := c.MultipartForm()
	for k := range form.File {
		fmt.Println(k)
	}
	fileheader := form.File["editormd-image-file"][0]
	fmt.Println(fileheader.Size)
	file, err := fileheader.Open()
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/error&code=403")
	}
	response, err := service.UploadToSMMS(generateSaveFile(fileheader.Filename, service.GenerateUUID()), file)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/error&code=403")
	}
	fmt.Println(response)
	if response.Success {
		c.JSON(http.StatusOK, gin.H{
			"success": 1,
			"message": response.Message,
			"url": response.Data.URL,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": 0,
			"message": "something wrong",
			"url": response.Data.URL,
		})
	}
}
// LoginController 登陆controller
func LoginController(c *gin.Context) {
	email := c.PostForm("email")
	passwd := c.GetString("password")
	author := service.GetAuthorByEmailAndPasswd(email, passwd)
	
	
	if author.Email != email {
		c.Redirect(http.StatusMovedPermanently, "/error&code=403")
	} else {
		var authorbs, _ = json.Marshal(author)

		session := sessions.Default(c)
		session.Set("author", string(authorbs))
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/home")
	}
}
// LogoutController 登出
func LogoutController(c *gin.Context) {

	session := sessions.Default(c)
	session.Delete("author")
	session.Save()

	fmt.Println("logout")

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
// RegisteController 负责处理注册的用户信息
func RegisteController(c *gin.Context) {
	form, _ := c.MultipartForm()
	file := form.File["picture"][0]
	picture := service.GenerateUUID()
	dst := generateSaveFile(file.Filename, picture)
	

	author := model.Author{
		Nick: c.PostForm("nickname"),
		Passwd: c.GetString("password"),
		Email: c.PostForm("email"),
		Description: c.PostForm("description"),
		Picture: dst,
	}
	fmt.Println(author)
	var id = service.RegisteAuthor(author)
	if id < 0 {
		c.Redirect(http.StatusMovedPermanently, "/error?code=403")
	} else {
		c.SaveUploadedFile(file, dst)


		//把author保存在session中
		var authorbs, _ = json.Marshal(author)

		session := sessions.Default(c)
		session.Set("author", string(authorbs))
		session.Save()

		c.Redirect(http.StatusMovedPermanently, "/home")
	}
}
func generateSaveFile(filename string, uuid string) string{
	return "picture/" + uuid + "." + strings.Split(filename, ".")[1]
}
// LostPasswordController 控制器
func LostPasswordController(c *gin.Context) {
	nickname := c.PostForm("nickname")
	email := c.PostForm("email")
	passwd := c.GetString("password")
	repasswd := c.GetString("repassword")
	if passwd != repasswd {
		c.Redirect(http.StatusMovedPermanently, "/lostpasswd")
		return
	} 
	if !service.VertifyAuthorByNicknameAndEmail(nickname, email) {
		c.Redirect(http.StatusMovedPermanently, "/lostpasswd")
		return
	}
	id := service.UpdatePasswordByEmail(email, passwd)
	if id < 0 {
		c.Redirect(http.StatusMovedPermanently, "/error?code=403")
	} else {
		author := service.GetAuthorByEmailAndPasswd(email, passwd)

		//把author保存在session中
		var authorbs, _ = json.Marshal(author)

		session := sessions.Default(c)
		session.Set("author", string(authorbs))
		session.Save()

		c.Redirect(http.StatusMovedPermanently, "/home")
	}
}