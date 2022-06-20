package linkFactory

import (
	"fmt"
	"github.com/go-redis/redis"
	"linkr-frame/app/constDir"
	"linkr-frame/global"

	//"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"sync"
)

type Pool interface {
	GetInstance() *Connect
	InitConnectPool() bool
	GetConnectLibrary() (interface{}, error)
}

var (
	once     sync.Once
	instance *Connect
	errDb    error
	db       *gorm.DB
	pool     *redis.Client
	redisDb  int
	dbType   string
)

type Connect struct {
	library string
}

func (this *Connect) GetInstance() *Connect {
	once.Do(func() {
		instance = &Connect{
			library: this.library,
		}
	})

	return instance
}

func (this *Connect) InitConnectPool() (result bool) {
	switch dbType {
	case "mysql":
		source := getDbLibrary(this.library)
		db, errDb = gorm.Open(
			mysql.Open(source), &gorm.Config{})
		if errDb != nil {
			log.Fatal(errDb.Error())
			return false
		}
		//链接池配置、集群数据源链接配置
		MaxIdleConns, _ := strconv.Atoi(global.Config.Mysql.MaxIdleCoons)
		MaxOpenConns, _ := strconv.Atoi(global.Config.Mysql.MaxOpenCoons)
		//增加sql配置
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(MaxOpenConns)
		// 设置最大空闲数
		sqlDB.SetMaxIdleConns(MaxIdleConns)
	case "redis":
		redisDb, _ := strconv.Atoi(global.Config.Redis.Db)
		pool = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", global.Config.Redis.Host, global.Config.Redis.Port), // redis地址
			Password: global.Config.Redis.Auth,                                                 // redis密码，没有则留空
			DB:       redisDb,                                                                  // 默认数据库，默认是0
		})

		//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
		_, err := pool.Ping().Result()
		if err != nil {
			log.Fatalf("redis链接异常：%s", err)
			return false
		}
	}
	return true
}

func (this *Connect) GetConnectLibrary() (res interface{}, err error) {

	switch dbType {
	case "mysql":
		return db, err
	case "redis":
		return pool, err
	default:
		return db, err
	}
}

func NewConnect(connect string, library string) *Connect {
	dbType = connect
	instance = &Connect{
		library: library,
	}

	return instance
}

//getDbLibrary 获取db dsn
func getDbLibrary(library string) string {
	var sourceMap = map[string]string{
		constDir.LinkIt: global.Config.Mysql.LinkItDb,
		constDir.Schema: global.Config.Mysql.SchemaDb,
	}
	source := sourceMap[library]
	source += "?charset=" + global.Config.Mysql.Charset +
		"&parseTime=True&loc=Local&timeout=5000ms"
	return source
}
