package router

import (
	"github.com/gin-gonic/gin"
	"music-saas/api"
)

func InitMusicRouter(Router *gin.RouterGroup) {
	MusicRouter := Router.Group("music")
	{
		MusicRouter.GET("id", api.GetMusicById) // 获取音乐详情
		MusicRouter.GET("", api.GetMusic)       // 音乐任务列表
		MusicRouter.POST("", api.CreateMusic)   // 增加音乐任务
		MusicRouter.PUT("", api.UpdateMusic)    // 编辑音乐任务
		MusicRouter.DELETE("", api.DeleteMusic) // 删除音乐任务
	}
}
