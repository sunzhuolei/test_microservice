package database

import (
	"closer_user/config"
	"closer_user/internal/pkg/global"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

/**
初始化数据库
 */
func InitDataBase()bool{
	configure := config.GetConfig()
	args := configure.Data.DataBase.Source
	driver := configure.Data.DataBase.Driver
	db, err := gorm.Open(driver, args+"?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai")
	if err != nil{
		fmt.Println("链接数据库失败：",err.Error())
		return false
	}
	db.DB().SetMaxOpenConns(1000)
	db.DB().SetMaxIdleConns(50)
	db.DB().SetConnMaxLifetime(time.Minute)
	db.DB().SetConnMaxIdleTime(time.Second * 30)
	log.Println("set mysql config success")
	global.Db = db
	global.Db.SingularTable(true) //默认不使用复数
	global.Db.LogMode(true)       //打印sql日志
	return true
}