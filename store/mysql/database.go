package mysql

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var RedisDB redis.Conn

func GetDB() *gorm.DB {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/hongbao?charset=utf8mb4&parseTime=True&loc=Local"
	DB,err =gorm.Open("mysql",dsn)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&EnvelopeDo{}, &Trade{}, &AccountDo{})
	return DB
}

func GetRedis() {
	var err error
	RedisDB, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")

	//defer RedisDB.Close()
}
//func GetRedis() {
//	// 实例化一个redis客户端
//	redisdb = redis.NewClient(&redis.Options{
//		Addr:      "localhost:6379", // ip:port
//		Password: "",				   // redis连接密码
//		DB:       10,					   // 选择的redis库
//		PoolSize: 20,						   // 设置连接数,默认是10个连接
//	})
//	if redisdb!=nil{
//		fmt.Println("连接成功")
//
//	}
//
//}

