package session

import (
	"beego-hananoq-blog/conf"
	"github.com/astaxie/beego/logs"
	"gopkg.in/boj/redistore.v1"
	"net/http"
)

var store *redistore.RediStore

func init() {
	redis := conf.GetConfig().Redis.Default
	var err error
	store, err = redistore.NewRediStore(redis.Db,
		"tcp",
		redis.Addr,
		"",
		[]byte(conf.GetConfig().Session.SecretKey))
	if err != nil {
		logs.Error("Init redis error:", err)
	}
}

func Get(r *http.Request, name string) (val interface{}, err error) {
	session, err := store.Get(r, conf.GetConfig().Session.Name)
	if err != nil {
		return
	}
	val = session.Values[name]
	return
}

func Set(w http.ResponseWriter, r *http.Request, name string, val interface{}) (err error) {
	session, err := store.Get(r, conf.GetConfig().Session.Name)
	if err != nil {
		return
	}
	session.Values[name] = val
	err = session.Save(r, w)
	return
}

func Delete(w http.ResponseWriter, r *http.Request, name string) (err error) {
	session, err := store.Get(r, conf.GetConfig().Session.Name)
	if err != nil {
		return
	}
	delete(session.Values, name)
	err = session.Save(r, w)
	return
}
