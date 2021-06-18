/**
  @author:panliang
  @data:2021/6/18
  @note

**/
package model

import (
	"fmt"
	"go_im/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

type BaseModel struct {
	ID uint64
}


// 初始化 grom
func ConnectDB() *gorm.DB {

	var err error

	var(
		host     = config.GetString("database.mysql.host")
		port     = config.GetString("database.mysql.port")
		database = config.GetString("database.mysql.database")
		username = config.GetString("database.mysql.username")
		password = config.GetString("database.mysql.password")
		charset  = config.GetString("database.mysql.charset")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")

	config := mysql.New(mysql.Config{
		DSN:dsn,
	})
	DB,err = gorm.Open(config,&gorm.Config{

		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	log.Println(err)

	return DB

}

