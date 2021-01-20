package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age int `json:"age" binding:"required,ageValid"`
}

func ageValid(field validator.FieldLevel) bool {
	return field.Field().Interface().(int) > 18
}

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("login", func(context *gin.Context) {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			_ = v.RegisterValidation("ageValid", ageValid)
		}
		var user Login
		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if user.Username != "michael" || user.Password != "123456" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"status": "SUCCESS",
		})
	})
	router.POST("upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		_ = context.SaveUploadedFile(file, "./")
		context.JSON(http.StatusOK, gin.H{
			"code": "SUCCESS",
			"data": file.Filename,
		})
	})
	v1 := router.Group("v1")
	{
		v1.GET("", func(context *gin.Context) {
			context.String(http.StatusOK, "hello world")
		})
		v1.GET(":name", func(context *gin.Context) {
			name := context.Param("name")
			context.String(http.StatusOK, "hello %s", name)
		})
	}
	v2 := router.Group("v2")
	{
		v2.POST("", func(context *gin.Context) {
			name := context.PostForm("name")
			message := context.PostForm("message")
			data := context.DefaultPostForm("data", "This is default data")
			context.JSON(http.StatusOK, gin.H{
				"code":    "SUCCESS",
				"message": message,
				"data":    data,
				"name":    name,
			})
		})
	}

	return router
}
