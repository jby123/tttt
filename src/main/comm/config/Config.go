package config

import (
	"fmt"
	"goAdmin/src/main/comm/cache"
	"goAdmin/src/main/comm/config/comm"
	"goAdmin/src/main/comm/database"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const DefaultDevelopmentEnv = "dev"

//总配置
type Conf struct {
	ServerConf Server `yaml:"Server"`
	EnvConf    Gin    `yaml:"Gin"`
}

//系统服务配置
type Server struct {
	Port int `yaml:"Port"`
}

type Gin struct {
	ActiveConf Profiles `yaml:"Profiles"`
}

//环境  dev/test/prod
type Profiles struct {
	Active string `yaml:"Active"`
}

var Config Conf

var RedisConf cache.RedisConfig

var DbConf database.DbConfig

var CommConfig comm.Conf

// 初始化解析参数
var Path string

var activePath string

func PreInit(relativePath string) {
	Path = relativePath + "/bootstrap.yaml"
}

func InitActivePath(relativePath, activeEnv string) {
	if len(activeEnv) <= 0 {
		activeEnv = DefaultDevelopmentEnv
	}
	activePath = relativePath + "/application_" + activeEnv + ".yaml"

}

// 1.从配置文件中加载配置
func InitCommConfig(relativePath string) error {
	PreInit(relativePath)
	content, err := ioutil.ReadFile(Path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	err = yaml.Unmarshal(content, &Config)
	InitActivePath(relativePath, Config.EnvConf.ActiveConf.Active)
	//初始化缓存配置信息
	InitRedisConfig(relativePath)
	//初始化db配置信息
	InitDbConfig(relativePath)
	//初始化通用配置信息
	InitBaseConfig(relativePath)
	return err
}
func InitRedisConfig(relativePath string) error {
	if len(Path) == 0 {
		InitCommConfig(relativePath)
	}
	content, err := ioutil.ReadFile(activePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	err = yaml.Unmarshal(content, &RedisConf)
	return err
}

func InitDbConfig(relativePath string) error {
	if len(Path) == 0 {
		InitCommConfig(relativePath)
	}
	content, err := ioutil.ReadFile(activePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	err = yaml.Unmarshal(content, &DbConf)
	return err
}

func InitBaseConfig(relativePath string) error {
	if len(Path) == 0 {
		InitCommConfig(relativePath)
	}
	content, err := ioutil.ReadFile(activePath)
	if err == nil {
		err = yaml.Unmarshal(content, &CommConfig)
	}
	return err
}
