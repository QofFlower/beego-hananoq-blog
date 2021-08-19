package conf

import (
	"github.com/astaxie/beego"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Session struct {
		Name      string `yaml:"name"`
		SecretKey string `yaml:"secret_key"`
	} `yaml:"session"`
	DB struct {
		Default DB
	}
	Redis struct {
		Default Redis `yaml:"default"`
		Request Redis `yaml:"request"`
	}
	OAuth2 struct {
		Client []Client `yaml:"client"`
	} `yaml:"oauth2"`
	ReleaseRouter []string `yaml:"release_router"`
	OSS           struct {
		Endpoint        string `yaml:"endpoint" json:"endpoint"`
		Region          string `yaml:"region" json:"region"`
		AccessKeyId     string `yaml:"accessKeyId" json:"accessKeyId"`
		AccessKeySecret string `yaml:"accessKeySecret" json:"accessKeySecret"`
		Bucket          string `yaml:"bucket" json:"bucket"`
	} `yaml:"OSS"`
	Sign struct {
		AppId  string `yaml:"app_id"`
		Nonce  string `yaml:"nonce"`
		Secret string `yaml:"secret"`
	} `yaml:"sign"`
}

type DB struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Client struct {
	Id     string  `yaml:"id"`
	Secret string  `yaml:"secret"`
	Name   string  `yaml:"name"`
	Domain string  `yaml:"domain"`
	Scope  []Scope `yaml:"scope"`
}

type Scope struct {
	Id    string `yaml:"id"`
	Title string `yaml:"title"`
}

var config *Config

func init() {
	runMode := beego.AppConfig.DefaultString("RunMode", "dev")
	configFilePath := "conf/config_" + runMode + ".yaml"
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(content, &config); err != nil {
		log.Fatalf("unmarshal yaml error: %v", err)
	}
}

func GetConfig() *Config {
	return config
}
