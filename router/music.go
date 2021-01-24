package router

import (
	"github.com/gin-gonic/gin"
	"music-saas/api"
)

func InitMusicRouter(Router *gin.RouterGroup) {
	MusicProductionRouter := Router.Group("music")
	{
		MusicProductionRouter.POST("production/list", api.MusicProductionList)       // 音乐任务列表
		MusicProductionRouter.POST("production/create", api.CreateMusicProduction)   // 增加音乐任务
		MusicProductionRouter.POST("production/update", api.UpdateMusicProduction)   // 编辑音乐任务
		MusicProductionRouter.DELETE("production/delete", api.DeleteMusicProduction) // 删除音乐任务
	}
}
