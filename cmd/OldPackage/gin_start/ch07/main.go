package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

// 定义一个全局翻译
var trans ut.Translator

type LoginForm struct {
	User     string `json:"user"  binding:"required,min=3,max=20"` // 多个用逗号隔开
	Password string `json:"password"  binding:"required"`
}

type SignupForm struct {
	Age        uint8  `json:"age"  binding:"gte=1,lte=130"` //gte 大于等于; lte 小于等于
	Name       string `json:"name"  binding:"required,min=3,max=20"`
	Email      string `json:"email"  binding:"required,email"`
	Password   string `json:"password"  binding:"required"`
	RePassword string `json:"re_password"  binding:"required,eqfield=Password"` // 跨字段了,用eqfield
}

func InitTrans(locale string) (err error) {
	// 修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用的语言坏境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) err", locale)
		}
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func main() {
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var loginform LoginForm
		if err := c.ShouldBind(&loginform); err != nil {
			// 把 err 类型转换一下
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"message": err.Error(),
				})
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs.Translate(trans),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
		})
	})

	router.POST("/Signup", func(c *gin.Context) {
		var signupform SignupForm
		if err := c.ShouldBind(&signupform); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
		})
	})
	_ = router.Run(":8083")
}
