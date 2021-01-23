package router

import (
	"github.com/gin-gonic/gin"
	"music-saas/api"
)

func InitMusicRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("song")
	{
		UserRouter.POST("list", api.SongList)       // 音乐任务列表
		UserRouter.POST("create", api.CreateSong)   // 增加音乐任务
		UserRouter.POST("update", api.UpdateSong)   // 编辑音乐任务
		UserRouter.DELETE("delete", api.DeleteSong) // 删除音乐任务
	}
}
