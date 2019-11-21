package config

import (
	cache2 "goAdmin/src/main/comm/cache"
	"goAdmin/src/main/comm/database"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)


//总配置
type Conf struct {
	ServerConf Server `yaml:"Server"`
	EnvConf    Iris `yaml:"Iris"`
}

//系统服务配置
type Server struct {
	Port int `yaml:"Port"`
}


type Iris struct {
	ActiveConf Profiles `yaml:"Profiles"`
}
//环境  dev/test/prod
type Profiles struct {
	Active string `yaml:"Active"`
}

var Config Conf

var RedisConf cache2.RedisConfig

var DbConf database.DbConfig

// 初始化解析参数
var Path string

var activePath string

//默认环境
var DefaultEnvActive string ="dev"

func init() {
	Path = "./src/resources/bootstrap.yaml"
}

func InitActivePath(activeEnv string) {
	if len(activePath) <=0 {
		activeEnv = DefaultEnvActive
	}
	activePath = "./src/resources/application_"+activeEnv+".yaml"

}


// 1.从配置文件中加载配置
func InitCommConfig() error {
	content, err := ioutil.ReadFile(Path)
	if err == nil {
		err = yaml.Unmarshal(content, &Config)
	}
	InitActivePath(Config.EnvConf.ActiveConf.Active)

	InitRedisConfig()
	InitDbConfig()
	return err
}
func InitRedisConfig() error {
	if len(Path) == 0 {
		InitCommConfig()
	}
	content, err := ioutil.ReadFile(activePath)
	if err == nil {
		err = yaml.Unmarshal(content, &RedisConf)
	}
	return err
}

func InitDbConfig() error {
	if len(Path) == 0 {
		InitCommConfig()
	}
	content, err := ioutil.ReadFile(activePath)
	if err == nil {
		err = yaml.Unmarshal(content, &DbConf)
	}
	return err
}


