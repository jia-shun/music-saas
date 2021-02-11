package utils

import (
	"music-saas/model"
	"music-saas/model/transfer"
	"time"
)

func TransferToMusic(musicInfo transfer.MusicInfo) (music model.Music) {
	var timeLayoutStr = "2006-01-02T15:04:05.000Z"
	music.ID = musicInfo.ID
	music.MusicName = musicInfo.MusicName
	music.CustomerName = musicInfo.CustomerName
	music.FinishStatus = musicInfo.FinishStatus
	music.UserID = musicInfo.UserID
	music.Price = musicInfo.Price
	music.PayStatus = musicInfo.PayStatus
	beganAt, _ := time.Parse(timeLayoutStr, musicInfo.BeganAt)
	music.BeganAt = beganAt
	finishedAt, _ := time.Parse(timeLayoutStr, musicInfo.FinishedAt)
	music.FinishedAt = finishedAt
	return music
}

func TransferToMusicInfo(music model.Music) (musicInfo transfer.MusicInfo) {
	var timeLayoutStr = "2006-01-02 15:04:05"
	musicInfo.ID = music.ID
	musicInfo.MusicName = music.MusicName
	musicInfo.CustomerName = music.CustomerName
	musicInfo.Price = music.Price
	musicInfo.PayStatus = music.PayStatus
	musicInfo.UserID = music.UserID
	musicInfo.BeganAt = music.BeganAt.Format(timeLayoutStr)
	musicInfo.FinishedAt = music.FinishedAt.Format(timeLayoutStr)
	musicInfo.FinishStatus = music.FinishStatus
	return musicInfo
}
