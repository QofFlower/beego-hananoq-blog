package models

import (
	"beego-hananoq-blog/conf"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	config := conf.GetConfig()
	host := config.DB.Default.Host
	port := config.DB.Default.Port
	username := config.DB.Default.User
	password := config.DB.Default.Password
	dbname := config.DB.Default.Dbname
	dbType := config.DB.Default.Type

	addr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable connect_timeout=5 search_path=%s",
		username, password, dbname, host, port, dbname)

	orm.RegisterDriver(dbType, orm.DRPostgres)
	if err := orm.RegisterDataBase("default", dbType, addr); err != nil {
		panic("Database register err:" + err.Error())
	}
}
