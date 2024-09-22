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
type User struct {
	ID           uint           // Standard field for the primary key
	MyName       string         `gorm:"colume:name"`
	Email        *string        // 一个指向字符串的指针, allowing for null values
	Age          uint8          // 一个未签名的8位整数
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // 创建时间（由GORM自动管理）
	UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
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

	// string条件 通过where查询
	var user User
	db.Where("name = ?", "jinzhu").First(&user) // 不区分大小写,要记住数据库的列名

	// Struct条件 会忽略查询时候的零值 0，使用下面的map就可以解决
	db.Where(&User{MyName: "jinzhu"}).First(&user) // 可以屏蔽细节

	var users []User
	// Map条件
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 18}).Find(&users)
	for _, user := range users {
		fmt.Println(user.ID)
	}

	/*
		查询方式条件有三种1.string 2.struct 3.map
	*/

}
