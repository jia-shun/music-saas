package initialize

import (
	"github.com/gin-gonic/gin"
	"music/global"
	"music/middleware"
	"music/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	//跨域
	Router.Use(middleware.Cors())
	global.LOG.Info("use middleware cors")

	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		router.InitMusicRouter(PrivateGroup)
	}
	global.LOG.Info("router register success")
	return Router
}