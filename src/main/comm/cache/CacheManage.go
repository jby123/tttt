package cache

import (
	"github.com/go-redis/redis"
	"goAdmin/src/main/comm/exception"
	"time"
)

func errHandler(err error) {
	if err != nil {
		if err == redis.Nil {
			return
		}
		exception.CacheException(err)
	}

}

//设置缓存
func Set(key, val string, ttl time.Duration) error {

	error := redisClient.Set(key, val, ttl).Err()
	errHandler(error)
	return error
}

func Get(key string) (string, error) {
	value, error := redisClient.Get(key).Result()
	errHandler(error)
	if error == redis.Nil {
		return value, nil
	}
	return value, error

}

func Exists(key string) (bool, error) {
	value, error := redisClient.Exists(key).Result()
	errHandler(error)

	if error == redis.Nil {
		return false, nil
	}
	if value > 0 {
		return true, nil
	}
	return false, nil

}

func Delete(key string) (bool, error) {
	value, error := redisClient.Del(key).Result()
	errHandler(error)
	if error == redis.Nil {
		return false, nil
	}
	if value >= 0 {
		return true, nil
	}
	return false, nil
}
