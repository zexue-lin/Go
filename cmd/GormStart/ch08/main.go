package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// 模型定义
type NewUser struct {
	ID           uint           // Standard field for the primary key
	MyName       string         `gorm:"colume:name"`
	Email        *string        // 一个指向字符串的指针, allowing for null values
	Age          uint8          // 一个未签名的8位整数
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // 创建时间（由GORM自动管理）
	UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
	DeletedAt    gorm.DeletedAt // 增加软删除字段，开启软删除
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

	//db.AutoMigrate(&NewUser{})
	//var users = []NewUser{{MyName: "jinzhu1"}, {MyName: "jinzhu2"}, {MyName: "jinzhu3"}}
	//db.Create(&users)

	// 软删除
	db.Delete(&NewUser{}, 1) // 执行的是update操作，会更新删除时间那一列

	var users []NewUser
	db.Find(&users)
	for _, user := range users {
		fmt.Println(user.MyName) // 被软删除的字段不会被查询出来
	}

	// 真实删除
	db.Unscoped().Delete(&NewUser{}, 2)
}
