package service
import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"crypto/md5"
	"encoding/hex"
	"mime/multipart"
	"bytes"
	"io"
	"net/http"
	"io/ioutil"
	"blog/model"
	"encoding/json"
	"os"
)
// GenerateUUID 生成UUID
func GenerateUUID() string {
	u1 := uuid.Must(uuid.NewV4(), nil)
	return u1.String()
}
// EncodePassword 加密密码
func EncodePassword(data string) string {
    h := md5.New()
    h.Write([]byte(data))
    return hex.EncodeToString(h.Sum(nil))
}
// UploadToSMMS 向SMMS云传图片并返回链接
func UploadToSMMS(filename string, file multipart.File) (responseDao model.ResponseDAO, err error) {
	defer file.Close()

    //创建一个模拟的form中的一个选项,这个form项现在是空的
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作, 设置文件的上传参数叫uploadfile, 文件名是filename,
	//相当于现在还没选择文件, form项里选择文件的选项
	fileWriter, err := bodyWriter.CreateFormFile("smfile", filename)
	if err != nil {
		return
	}

	//iocopy 这里相当于选择了文件,将文件放到form中
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return
	}

    //获取上传文件的类型,multipart/form-data; boundary=...
	contentType := bodyWriter.FormDataContentType()


	//这个很关键,必须这样写关闭,不能使用defer关闭,不然会导致错误
	bodyWriter.Close()


    //这里就是上传的其他参数设置,可以使用 bodyWriter.WriteField(key, val) 方法
    //也可以自己在重新使用  multipart.NewWriter 重新建立一项,这个再server 会有例子
	params := map[string]string{
        "format" : "json",
    }
	//这种设置值得仿佛 和下面再从新创建一个的一样
	for key, val := range params {
		_ = bodyWriter.WriteField(key, val)
	}
	

	//发送post请求到服务端
	req, _ := http.NewRequest("POST", "https://sm.ms/api/v2/upload", bodyBuf)
	req.Header.Add("Authorization","cXf7hNRn4hxC2eEYqtWFOiyaVDyrQcSy")
	req.Header.Add("Content-Type", contentType)
	client := &http.Client{}
	resp, err := client.Do(req)
	// resp, err := http.Post("https://sm.ms/api/v2/upload", contentType, bodyBuf)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	responseDao = model.ResponseDAO{}
	err = json.Unmarshal(respbody, &responseDao)
	if err != nil {
		return
	}
	return
}
// SaveArticle 保存文章
func SaveArticle(content string) (string, error) {
	uuid := GenerateUUID()
	filename := "./article/" + uuid + ".md"
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	_, err = file.WriteString(content)
	defer file.Close()
	if err != nil {
		return "", err
	}
	return filename, nil
}
// GetAuthorBySession 通过session获取author
func GetAuthorBySession(c *gin.Context) model.Author{
	session := sessions.Default(c)
	var author model.Author
	if session.Get("author") != nil {
		// author = session.Get("author").(model.Author)
		// fmt.Printf("%T", session.Get("author"))
		json.Unmarshal([]byte(session.Get("author").(string)), &author)
	}
	return author
}