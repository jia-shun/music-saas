package router

import (
	"github.com/gin-gonic/gin"
	"music/api"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", api.Login)
		BaseRouter.POST("register", api.Register)
	}
	return BaseRouter
}
