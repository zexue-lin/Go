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

// User 有多张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
	gorm.Model
	Number    string
	UserRefer uint
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

	errs := db.AutoMigrate(&User{}, &CreditCard{})
	if errs != nil {
		panic(errs)
	}

	// 大型的系统中，不建议使用外键约束，会影响性能，也有有点:数据的完整性
	/*
			外键约束会让数据很完整，
		高并发的系统中一般不添加外键约束，在业务层面保证数据的一致性
	*/

	// 插入用户
	//user := User{}
	//db.Create(&user)
	//db.Create(&CreditCard{
	//	Number:    "12",
	//	UserRefer: user.ID,
	//})
	//db.Create(&CreditCard{
	//	Number:    "23",
	//	UserRefer: user.ID,
	//})

	// 查询用户
	var user User
	db.Preload("CreditCards").First(&user)
	for _, card := range user.CreditCards {
		fmt.Println(card.Number)
	}
}
