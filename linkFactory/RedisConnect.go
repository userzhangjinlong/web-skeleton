package linkFactory

import (
	"github.com/go-redis/redis"
)

//NewRedis redis工厂加载redis连接池
func NewRedis() (result bool) {
	result = NewConnect("redis", "").GetInstance().InitConnectPool()

	return result
}

//SelectDb redis工厂选择加载的redis库
func SelectDb(db int) (result bool) {
	redisDb = db
	result = NewConnect("redis", "").GetInstance().InitConnectPool()

	return result
}

//GetRedis redis工厂获取redis连接池
func GetRedis() (redisPool *redis.Client, err error) {
	if !NewRedis() {
		panic("redis链接异常")
	}
	redisConnect, errRedis := NewConnect("redis", "").GetInstance().GetConnectLibrary()

	return redisConnect.(*redis.Client), errRedis
}
