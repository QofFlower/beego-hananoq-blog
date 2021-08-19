package utils

import (
	"beego-hananoq-blog/conf"
	"crypto/md5"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetClient(clientId string) (client conf.Client) {
	for _, v := range conf.GetConfig().OAuth2.Client {
		if clientId == v.Id {
			client = v
			return
		}
	}
	return
}

func GetScopes(clientId, scope string) (scopes []conf.Scope) {
	client := GetClient(clientId)
	params := strings.Split(scope, ",")
	for _, p := range params {
		for _, v := range client.Scope {
			if p == v.Id {
				scopes = append(scopes, v)
			}
		}
	}
	return
}

func ScopeJoin(scopes []conf.Scope) string {
	var s []string
	for _, scope := range scopes {
		s = append(s, scope.Id)
	}
	return strings.Join(s, ",")
}

// Md5Encode md5加密
func Md5Encode(param string) string {
	w := md5.New()
	if _, err := io.WriteString(w, param); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", w.Sum(nil))
}

// IsNum 判断字符串是否为数字
func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// UriFilter 判断url是否为放行的
func UriFilter(url string) bool {
	for _, router := range conf.GetConfig().ReleaseRouter {
		if strings.Contains(url, router) {
			return true
		}
	}
	return false
}

// RemoveDuplicate 数组去重
func RemoveDuplicate(s []string) (res []string) {
	sort.Strings(s)
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && s[i] == s[i+1] {
			continue
		}
		res = append(res, s[i])
	}
	return
}

func UnixToTime(e string) (dataTime time.Time, err error) {
	data, err := strconv.ParseInt(e, 10, 64)
	dataTime = time.Unix(data/1000, 0)
	return
}
