package cache


import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisClient *redis.Client


//redis常量
type RedisConfig struct {
	Cache  CacheConf `yaml:"Cache"`
}
type CacheConf struct {
	Network string `yaml:"NetWork"` //连接类型 TCP
	DB      int    `yaml:"DB"`
	Host    string `yaml:"Host"`
	Port    int    `yaml:"Post"`
	Pwd     string `yaml:"Pwd"`
	MaxIdle int    `yaml:"maxIdle"`
	MaxOpen int    `yaml:"maxOpen"`
}




// 初始化redis
func InitRedis(conf *RedisConfig) (err error) {
	redisClient = redis.NewClient(
		&redis.Options{
			Network:     conf.Cache.Network,
			Addr:        conf.Cache.Host + ":" + fmt.Sprintf("%d", conf.Cache.Port),
			Password:    conf.Cache.Pwd,
			DB:          conf.Cache.DB,
			IdleTimeout: time.Duration(30) * time.Minute,
		})

	redisClient.WrapProcess(func(old func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
		return func(cmd redis.Cmder) error {
			fmt.Printf("starting process:<%s>\n", cmd)
			err := old(cmd)
			fmt.Printf("finished process:<%s>\n", cmd)
			return err
		}
	})

	return
}

// 关闭redis  只有在资源回收系统关闭时处理
func CloseRedis() {
	errorMsg := recover()
	if errorMsg != nil {
		fmt.Printf("CloseRedis.receive.errorMsg:%s", errorMsg)
	}
	if redisClient != nil {
		redisClient.Close()
	}
}
