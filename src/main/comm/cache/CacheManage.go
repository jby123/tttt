package cache

import (
	"fmt"
	"time"
)

func errHandler(err error) {
	if err != nil {
		fmt.Printf("err_handler, error:%s\n", err.Error())
		panic(err.Error())
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
	if error != nil {
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
