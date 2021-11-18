/**
  @author:panliang
  @data:2021/6/18
  @note

**/
package model

import (
	"fmt"
	"im_app/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

type BaseModel struct {
	ID uint64
}

// 初始化 grom
func ConnectDB() *gorm.DB {
	var (
		host     = config.GetString("database.mysql.host")
		port     = config.GetString("database.mysql.port")
		database = config.GetString("database.mysql.database")
		username = config.GetString("database.mysql.username")
		password = config.GetString("database.mysql.password")
		charset  = config.GetString("database.mysql.charset")
		//loc  = config.GetString("database.mysql.loc")
		err error
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		username, password, host, port, database, charset)

	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
	//	username, password, host, port, database, charset, true, url.QueryEscape(loc))
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		fmt.Println("Mysql 连接异常: ")
		panic(err.Error())
	}
	return DB
}
