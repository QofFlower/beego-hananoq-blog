package conf

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
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
	if err := initConfigFromNacos(); err != nil {
		logs.Error(err)
	}
}

func GetConfig() *Config {
	return config
}

func initConfigFromFile() error {
	runMode := beego.AppConfig.DefaultString("RunMode", "dev")
	configFilePath := "conf/config_" + runMode + ".yaml"
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = parseConfig(content)
	return err
}

func initConfigFromNacos() error {
	var sc []constant.ServerConfig
	port, _ := beego.AppConfig.Int("nacos::port")
	sc = append(sc, constant.ServerConfig{
		IpAddr: beego.AppConfig.String("nacos::addr"),
		Port:   uint64(port),
	})
	namespace := beego.AppConfig.String("nacos::namespace")
	configDataID := beego.AppConfig.String("nacos::data_id")
	configGroup := beego.AppConfig.String("nacos::group")

	if len(sc) == 0 ||
		namespace == "" ||
		configDataID == "" ||
		configGroup == "" {
		return fmt.Errorf("nacos config invalid, namespace:%s configDataID:%s configGroup:%s sc:%+v",
			namespace, configDataID, configGroup, sc)
	}

	cc := constant.ClientConfig{
		NamespaceId:         namespace, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./logs/nacos",
		CacheDir:            "./docs/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// a more graceful way to create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return err
	}

	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: configDataID,
		Group:  configGroup,
	})
	if err != nil {
		return err
	}
	if err = parseConfig([]byte(content)); err != nil {
		return err
	}

	go func() {
		forever := make(chan struct{})

		err := client.ListenConfig(vo.ConfigParam{
			DataId: configDataID,
			Group:  configGroup,
			OnChange: func(namespace, group, dataId, data string) {
				fmt.Println("on config changed, group:" + group + ", dataId:" + dataId + ", content:" + data)
				if err := parseConfig([]byte(data)); err != nil {
					fmt.Println("on config changed, parse err:", err)
				}
			},
		})
		if err != nil {
			fmt.Println("nacos config listen err:", err)
		}

		_ = <-forever
	}()
	return nil
}

func parseConfig(content []byte) error {
	c := &Config{}
	if err := yaml.Unmarshal(content, c); err != nil {
		return err
	}
	//rwLock.Lock()
	config = c
	//rwLock.Unlock()
	return nil
}
