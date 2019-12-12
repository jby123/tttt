package cache

import (
	"goAdmin/src/main/comm/exception"
	"time"
)

func errHandler(err error) {
	if err != nil {
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
	if len(value) <= 0 {
		return "", error
	}
	return value, error

}

func Exists(key string) (bool, error) {
	value, error := redisClient.Exists(key).Result()
	errHandler(error)
	if error != nil {
		return false, error
	}
	if value >= 0 {
		return true, nil
	}
	return false, nil
}

func Delete(key string) (bool, error) {
	value, error := redisClient.Del(key).Result()
	errHandler(error)
	if error != nil {
		return false, error
	}
	if value >= 0 {
		return true, nil
	}
	return false, nil
}
