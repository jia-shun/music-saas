package router

import (
	"github.com/gin-gonic/gin"
	"music-saas/api"
)

func InitUserRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	BaseRouter := Router.Group("user")
	{
		BaseRouter.POST("info", api.GetInfo)
		BaseRouter.POST("logout", api.Logout)
	}
	return BaseRouter
}
