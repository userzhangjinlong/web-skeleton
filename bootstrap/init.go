package bootstrap

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"linkr-frame/app/constDir"
	"linkr-frame/config"
	"linkr-frame/global"
	"linkr-frame/linkFactory"
	"log"
	"os"
	"time"
)

var (
	System = new(config.System)
)

//加载全局系统配置
func InitSystemConfig() {
	dir, _ := os.Getwd()
	file, err := ioutil.ReadFile(dir + "/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(file, System)
	global.Config = System
	if err != nil {
		log.Fatal("初始化系统配置异常")
		return
	}
}

//初始化日志配置
func InitLogConfig() {
	//设置日志格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999999999",
	})
	dir, _ := os.Getwd()
	filePath := dir + global.Config.Log.Path

	//日志响应配置
	writer, _ := rotatelogs.New(
		filePath+"%Y%m%d%H%M"+".logs",
		rotatelogs.WithLinkName(filePath+"/logs"),
		rotatelogs.WithMaxAge(time.Hour*24*15),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithRotationSize(100*1024*1024),
	)
	logrus.SetOutput(writer)
}

//初始化数据库
func InitMysql() {

	db, err := linkFactory.GetMysql(constDir.LinkIt)
	if err != nil {
		panic(fmt.Sprintf("linkIt 数据库连接异常:%s", err))
	}
	schemaDb, err := linkFactory.GetMysql(constDir.Schema)
	if err != nil {
		panic(fmt.Sprintf("schemaDb 数据库连接异常:%s", err))
	}
	global.LinkItDB = db
	global.SchemaDB = schemaDb
}

//初始化redis
func InitRedis() {
	redis, err := linkFactory.GetRedis()
	if err != nil {
		panic(fmt.Sprintf("linkIt redis连接异常:%s", err))
	}
	global.Redis = redis
}
