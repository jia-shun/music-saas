package service

import (
	"errors"
	"music-saas/global"
	"music-saas/model"
	"music-saas/model/request"
	"time"
)

func GetMusicList(info request.PageInfo, keyword string, order string, desc bool, userId uint) (list []model.Music, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&model.Music{})
	var musicList []model.Music

	if keyword != "" {
		db = db.Where("`user_id` = ? AND (music_name LIKE ? OR customer_name LIKE ?)", userId, "%"+keyword+"%", "%"+keyword+"%")
	} else {
		db = db.Where("`user_id` = ?", userId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return musicList, total, err
	}
	db = db.Limit(limit).Offset(offset)
	if order != "" {
		var OrderStr string
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		}
		err = db.Order(OrderStr).Find(&musicList).Error
	} else {
		err = db.Order("updated_at desc").Find(&musicList).Error
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
