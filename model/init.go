package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(connstring string) {
	fmt.Println("mysql数据库连接中...", connstring)
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		fmt.Println("错误信息：", err)
		panic("mysql数据库连接错误")

	}
	fmt.Println("mysql数据库连接成功")
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)       //表名不加s,users->user
	db.DB().SetMaxIdleConns(20)  //设置连接池
	db.DB().SetMaxOpenConns(100) //最大连接数
	db.DB().SetConnMaxLifetime(30 * time.Second)
	DB = db
	migration()
}
