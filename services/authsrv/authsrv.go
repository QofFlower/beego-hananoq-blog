package authsrv

import (
	"beego-hananoq-blog/models"
	"crypto/md5"
	"fmt"
	"io"
)

func GetUserIdByUsernameAndPwd(username, password string) (userID string) {
	username, password = Md5Encode(username), Md5Encode(password)
	return models.GetUserIdByUsernameAndPwd(username, password)
}

func Md5Encode(param string) (encode string) {
	w := md5.New()
	if _, err := io.WriteString(w, param); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", w.Sum(nil))
}
