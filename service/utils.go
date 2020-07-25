package service
import (
	"github.com/satori/go.uuid"
	"crypto/md5"
	"encoding/hex"
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