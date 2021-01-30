package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"music-saas/global"
	"music-saas/model"
	"music-saas/model/request"
	"music-saas/model/response"
	"music-saas/service"
	"music-saas/utils"
)

func GetMusic(ctx *gin.Context) {
	var pageInfo request.SearchMusicParams
	_ = ctx.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if list, total, err := service.GetMusicList(pageInfo.Keyword, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
		global.LOG.Error("获取音乐列表失败", zap.Any("err", err))
		response.FailWithMessage("获取音乐列表失败", ctx)
	} else {
		response.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, ctx)
	}
}

func GetMusicById(ctx *gin.Context) {
	var music model.Music
	_ = ctx.ShouldBindJSON(&music)
	if err := utils.Verify(music.MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	music, err := service.FindMusicById(music.ID)
	if err != nil {
		global.LOG.Error("获取音乐失败", zap.Any("err", err))
		response.FailWithMessage("获取音乐失败", ctx)
	} else {
		response.OkWithData(music, ctx)
	}
}

func CreateMusic(ctx *gin.Context) {
	var music model.Music
	_ = ctx.ShouldBindJSON(&music)
	if err := utils.Verify(music, utils.MusicVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := service.CreateMusic(music); err != nil {
		global.LOG.Error("创建音乐失败", zap.Any("err", err))
		response.FailWithMessage("创建音乐失败", ctx)
	} else {
		response.Ok(ctx)
	}
}

func UpdateMusic(ctx *gin.Context) {
	var music model.Music
	_ = ctx.ShouldBindJSON(&music)
	if err := utils.Verify(music.MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := utils.Verify(music, utils.MusicVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := service.UpdateMusic(music); err != nil {
		global.LOG.Error("编辑音乐失败", zap.Any("err", err))
		response.FailWithMessage("编辑音乐失败", ctx)
	} else {
		response.Ok(ctx)
	}
}

func DeleteMusic(ctx *gin.Context) {
	var music model.Music
	_ = ctx.ShouldBindJSON(&music)
	if err := utils.Verify(music.MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := service.DeleteMusic(music); err != nil {
		global.LOG.Error("删除音乐失败", zap.Any("err", err))
		response.FailWithMessage("删除音乐失败", ctx)
	} else {
		response.Ok(ctx)
	}
}
