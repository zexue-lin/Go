package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// gorm基础知识
// Product 表定义，其实就是定义一个 struct
type Product struct {
	gorm.Model // 匿名的功能性字段，可以直接使用
	Code       sql.NullString
	Price      uint
}

// loc=Local  是本地时区的意思
// 先用gorm打开mysql数据库，再可以添加一些配置 Config
func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	// 【日志配置】设置全局的logger，这个logger在执行每个sql语句的时候都会打印每一行sql
	// sql才是最重要的，std -> standard 标准的
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

	// 上面定义一个表结构。将表结构直接生成对应的表 - migrations
	// 迁移 schema
	_ = db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: sql.NullString{"", true}, Price: 100})

	// Read
	// db.First 读取第一条记录,传递的是一个指针（可以直接修改值）
	var product Product
	db.First(&product, 1)                 // 根据整型主键查找,查找id=1的
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{"", true}}) // !!! 仅更新非零值字段 !!!为空或为0不更新
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	// 并没有执行 deleted 语句，逻辑删除，并没有真实删除，（大公司内部很少使用物理删除）
	//db.Delete(&product, 1) // 通过主键删除
}
