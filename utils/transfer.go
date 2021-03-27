package utils

import (
	"music-saas/model"
	"music-saas/model/transfer"
	"strconv"
	"strings"
	"time"
)

var weekday = [7]string{"星期天", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}

const timeLayoutStr = "2006-01-02"
const timeParseStr = "2006-01-02T15:04:05.000Z"

func TransferToMusic(musicInfo transfer.MusicInfo) (music model.Music) {
	music.ID = musicInfo.ID
	music.MusicName = musicInfo.MusicName
	music.CustomerName = musicInfo.CustomerName
	music.FinishStatus = musicInfo.FinishStatus
	music.UserID = musicInfo.UserID
	music.Price = musicInfo.Price
	music.PayStatus = musicInfo.PayStatus
	beganAt, _ := time.Parse(timeParseStr, musicInfo.BeganAt)
	music.BeganAt = beganAt
	finishedAt, _ := time.Parse(timeParseStr, musicInfo.FinishedAt)
	music.FinishedAt = finishedAt
	return music
}

func TransferToMusicInfo(music model.Music) (musicInfo transfer.MusicInfo) {
	musicInfo.ID = music.ID
	musicInfo.MusicName = music.MusicName
	musicInfo.CustomerName = music.CustomerName
	musicInfo.Price = music.Price
	musicInfo.PayStatus = music.PayStatus
	musicInfo.UserID = music.UserID
	musicInfo.BeganAt = timeToViewString(music.CreatedAt)
	musicInfo.FinishedAt = timeToViewString(music.FinishedAt)
	musicInfo.FinishStatus = music.FinishStatus
	return musicInfo
}

func timeToViewString(timeAt time.Time) string {
	timeAtDay := timeAt.Format(timeLayoutStr)
	timeAtWeek := function2Week(timeAtDay)
	return strings.Join([]string{timeAtDay, timeAtWeek}, "-")
}

func function2Week(timeStr string) string {
	var year, month, day uint16
	timeSplit := strings.Split(timeStr, "-")
	intYearNum, _ := strconv.Atoi(timeSplit[0])
	year = uint16(intYearNum)
	intMonthNum, _ := strconv.Atoi(timeSplit[1])
	month = uint16(intMonthNum)
	intDayNum, _ := strconv.Atoi(timeSplit[2])
	day = uint16(intDayNum)
	var y, m, c uint16
	if month >= 3 {
		m = month
		y = year % 100
		c = year / 100
	} else {
		m = month + 12
		y = (year - 1) % 100
		c = (year - 1) / 100
	}

	week := y + (y / 4) + (c / 4) - 2*c + ((26 * (m + 1)) / 10) + day - 1
	if week < 0 {
		week = 7 - (-week)%7
	} else {
		week = week % 7
	}
	whichWeek := int(week)
	return weekday[whichWeek]
}
