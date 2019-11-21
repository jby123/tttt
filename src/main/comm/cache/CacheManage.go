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

	error := redisClient.Set(key,val,ttl).Err()
	errHandler(error)
	return error
}

func Get(key string) (string,error) {
	value, error := redisClient.Get(key).Result()
	errHandler(error)
	if error != nil {
		return "", error
	}
	return value, error

}