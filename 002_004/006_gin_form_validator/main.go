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
	"reflect"
	"strings"
)

type LoginForm struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required,min=3,max=30"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type SignUpInForm struct {
	Age        uint8  `form:"age" json:"age" xml:"age" binding:"gte=1,lte=130"`
	Name       string `form:"name" json:"name" xml:"name" binding:"required"`
	Email      string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Password   string `form:"password" json:"password" xml:"password" binding:"required"`
	RePassword string `form:"re_password" json:"re_password" xml:"re_password" binding:"required,eqfield=Password"`
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for k, v := range fileds {
		rsp[k[strings.Index(k, ".")+1:]] = v
	}
	return rsp
}

func InitTrans(local string) (err error) {
	// 修改gin框架中的validator的属性，定制返回错误信息
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册获取json的tag的自定义方法
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()
		// 第一个参数是备用语言环境，后面的是应该支持的语言
		uni := ut.New(enT, zhT, enT)
		trans, ok = uni.GetTranslator(local)
		if !ok {
			return fmt.Errorf("translator not found translator: %s", local)
		}
		switch local {
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(validate, trans)
		case "en":
			_ = en_translations.RegisterDefaultTranslations(validate, trans)
		default:
			_ = en_translations.RegisterDefaultTranslations(validate, trans)
		}
	}
	return
}

var trans ut.Translator

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if err := c.ShouldBind(&form); err != nil {
			// 自定义表单返回错误信息
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": errs})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": removeTopStruct(errs.Translate(trans)),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "ok",
			"data": form,
		})
	})
	router.POST("/signup", func(c *gin.Context) {
		var form SignUpInForm
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "ok",
			"user": form.Age,
		})
	})

	router.Run(":8080")
}
