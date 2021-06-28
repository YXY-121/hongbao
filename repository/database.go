package repository

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func getDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123@(127.0.0.1:3306)/hongbao?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	db.AutoMigrate(&Envelope{},&Trade{},&Account{},&User{})

	return db
}

func getRedis()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")

	defer c.Close()
}