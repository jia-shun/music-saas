package router

import (
	"github.com/gin-gonic/gin"
	"music-saas/api"
)

func InitMusicRouter(Router *gin.RouterGroup) {
	MusicRouter := Router.Group("music")
	{
		MusicRouter.GET(":id", api.GetMusicById)                 // 获取音乐详情
		MusicRouter.POST("list", api.GetMusic)                   // 音乐任务列表
		MusicRouter.POST("create", api.CreateMusic)              // 增加音乐任务
		MusicRouter.POST("update", api.UpdateMusic)              // 编辑音乐任务
		MusicRouter.POST("update-status", api.UpdateMusicStatus) // 编辑音乐任务状态
		MusicRouter.POST("delete", api.DeleteMusic)              // 删除音乐任务
	}
}
