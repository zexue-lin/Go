package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID int     // 数据库中存储的字段company_id
	Company   Company // 真正使用的是这个
}

type Company struct {
	ID   int
	Name string
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log 级别
			//IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			//ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful: true, // 彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	errs := db.AutoMigrate(&User{}) // 新建了user表和company表，并设置了外键
	if errs != nil {
		panic(errs)
	}

	var user User

	// 这里只查询了user表
	//db.First(&user)
	//fmt.Println(user.Name, user.Company.ID)

	// 1.Preload
	//db.Preload("Company").First(&user)
	//fmt.Println(user.Name, user.Company.Name)

	// 2.Joins
	db.Joins("Company").First(&user)
	fmt.Println(user.Name, user.Company.Name)
}
