package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// User 拥有并属于多种 language，`user_languages` 是连接表
type User3 struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	//gorm.Model
	Name    string
	AddTime time.Time // 每个记录插入的时候自动插入当前时间
}

func (l *Language) BeforeCreate(tx *gorm.DB) (err error) {
	l.AddTime = time.Now()
	return
}

func (Language) TableName() string {
	return "my_languages"
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
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "pre_",
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	//errs := db.AutoMigrate(&User3{})
	//if errs != nil {
	//	panic(errs)
	//}

	// 插入数据
	//Languages := []Language{}
	//Languages = append(Languages, Language{Name: "go"})
	//Languages = append(Languages, Language{Name: "java"})
	//user := User3{
	//	Languages: Languages,
	//}
	//db.Create(&user)

	// 查询数据1
	//var user User3
	//db.Preload("Languages").First(&user)
	//for _, lang := range user.Languages {
	//	fmt.Println(lang.Name) // go java
	//}

	// 查询数据2 关联模式
	var user User3
	db.First(&user)
	var lang []Language
	_ = db.Model(&user).Association("Languages").Find(&lang)
	for _, lang := range lang {
		fmt.Println(lang.Name) // go java
	}
}
