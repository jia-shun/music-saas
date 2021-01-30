package service

import (
	"music-saas/global"
	"music-saas/model"
	"music-saas/model/request"
)

func GetMusicList(keyword string, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&model.Music{})
	var musicList []model.Music
	if keyword != "" {
		db = db.Where("song_name LIKE ?", "%"+keyword+"%")
		db = db.Where("customer_name LIKE ?", "%"+keyword+"%")
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
		err = db.Order("updated_at").Find(&musicList).Error
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
	if oldMusic, err := FindMusicById(music.ID); err != nil {
		return err
	} else {
		oldMusic.SongName = music.SongName
		oldMusic.CustomerName = music.CustomerName
		oldMusic.Price = music.Price
		return global.DB.Save(&oldMusic).Error
	}
}

func DeleteMusic(music model.Music) (err error) {
	return global.DB.Delete(&music).Error
}
