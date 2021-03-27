package service

import (
	"errors"
	"music-saas/global"
	"music-saas/model"
	"music-saas/model/request"
	"time"
)

func GetMusicList(info request.SearchMusicParams, userId uint) (list []model.Music, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&model.Music{})
	var musicList []model.Music
	var keyword = info.Keyword
	var payStatusStr = info.PayStatus
	var finishStatusStr = info.FinishStatus
	var order = info.OrderKey
	var desc = info.Desc
	db = db.Where("`user_id` = ?", userId)
	if keyword != "" {
		db = db.Where("music_name LIKE ? OR customer_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if payStatusStr != "" {
		var payStatus int
		if "TRUE" == payStatusStr {
			payStatus = 1
		} else {
			payStatus = 0
		}
		db = db.Where("pay_status = ?", payStatus)
	}
	if finishStatusStr != "" {
		var finishStatus int
		if "TRUE" == finishStatusStr {
			finishStatus = 1
		} else {
			finishStatus = 0
		}
		db = db.Where("finish_status = ?", finishStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return musicList, total, err
	}
	db = db.Limit(limit).Offset(offset)
	if order != "" {
		var orderStr string
		if "beganAt" == order {
			orderStr = "began_at"
		}
		if "createdAt" == order {
			orderStr = "created_at"
		}
		if desc {
			orderStr = orderStr + " desc"
		}
		err = db.Order(orderStr).Find(&musicList).Error
	} else {
		err = db.Order("finished_at desc").Find(&musicList).Error
	}
	return musicList, total, err
}

func FindMusicById(id uint) (music model.Music, err error) {
	err = global.DB.Where("`id` = ?", id).First(&music).Error
	return music, err
}

func CreateMusic(music model.Music) (err error) {
	return global.DB.Create(&music).Error
}

func UpdateMusic(music model.Music) (err error) {
	oldMusic, err := FindMusicById(music.ID)
	if err != nil {
		return err
	}
	if oldMusic.UserID != music.UserID {
		return errors.New("没有权限修改这首音乐")
	}
	oldMusic.FinishStatus = music.FinishStatus
	oldMusic.MusicName = music.MusicName
	oldMusic.CustomerName = music.CustomerName
	oldMusic.Price = music.Price
	oldMusic.PayStatus = music.PayStatus
	return global.DB.Save(&oldMusic).Error
}

func UpdateMusicFinishStatus(music model.Music) (err error) {
	oldMusic, err := FindMusicById(music.ID)
	if err != nil {
		return err
	}
	if oldMusic.UserID != music.UserID {
		return errors.New("没有权限修改这首音乐")
	}
	oldMusic.FinishStatus = music.FinishStatus
	if oldMusic.FinishStatus {
		oldMusic.FinishedAt = time.Now()
	}
	return global.DB.Save(&oldMusic).Error
}

func UpdateMusicPayStatus(music model.Music) (err error) {
	oldMusic, err := FindMusicById(music.ID)
	if err != nil {
		return err
	}
	if oldMusic.UserID != music.UserID {
		return errors.New("没有权限修改这首音乐")
	}
	oldMusic.PayStatus = music.PayStatus
	return global.DB.Save(&oldMusic).Error
}

func DeleteMusic(music model.Music) (err error) {
	oldMusic, err := FindMusicById(music.ID)
	if err != nil {
		return err
	}
	if oldMusic.UserID != music.UserID {
		return errors.New("没有权限删除这首音乐")
	}
	return global.DB.Delete(&oldMusic).Error
}
