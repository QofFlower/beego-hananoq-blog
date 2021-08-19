package logger

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"log"
	"strings"
)

func InitLogger() (err error) {
	configure, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		log.Fatalf("Config init error: %v", err)
	}
	maxLines, err := configure.Int64("log::maxlines")
	if err != nil {
		maxLines = 100
	}

	appName := beego.AppConfig.DefaultString("AppName", "beego")
	logPath := configure.String("log::log_path")
	l, r := strings.IndexByte(logPath, '{'), strings.LastIndexByte(logPath, '}')
	var fileName string
	if l >= r {
		fileName = logPath
	} else {
		var strs []string
		bytes := []byte(logPath)
		strs = append(strs, string(bytes[:l]), appName, string(bytes[r+1:]))
		fileName = strings.Join(strs, "")
	}
	logConf := make(map[string]interface{})
	logConf["filename"] = fileName
	level, _ := configure.Int("log::log_level")
	logConf["level"] = level
	logConf["maxlines"] = maxLines

	confStr, err := json.Marshal(logConf)
	if err != nil {
		log.Fatalf("failed to marshal error: %v", err)
	}
	if err = logs.SetLogger(logs.AdapterFile, string(confStr)); err != nil {
		log.Fatalf("set logger error: %v", err)
		return
	}
	logs.SetLogFuncCall(true)
	return
}
