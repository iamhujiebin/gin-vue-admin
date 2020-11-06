package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateAdRewards
// @description   create a AdRewards
// @param     adRewards               model.AdRewards
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAdRewards(adRewards model.AdRewards) (err error) {
	err = global.GVA_DB_LOGIC.Create(&adRewards).Error
	return err
}

// @title    DeleteAdRewards
// @description   delete a AdRewards
// @auth                     （2020/04/05  20:22）
// @param     adRewards               model.AdRewards
// @return                    error

func DeleteAdRewards(adRewards model.AdRewards) (err error) {
	err = global.GVA_DB_LOGIC.Delete(adRewards).Error
	return err
}

// @title    DeleteAdRewardsByIds
// @description   delete AdRewardss
// @auth                     （2020/04/05  20:22）
// @param     adRewards               model.AdRewards
// @return                    error

func DeleteAdRewardsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB_LOGIC.Delete(&[]model.AdRewards{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateAdRewards
// @description   update a AdRewards
// @param     adRewards          *model.AdRewards
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAdRewards(adRewards *model.AdRewards) (err error) {
	err = global.GVA_DB_LOGIC.Save(adRewards).Error
	return err
}

// @title    GetAdRewards
// @description   get the info of a AdRewards
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    AdRewards        AdRewards

func GetAdRewards(id uint) (err error, adRewards model.AdRewards) {
	err = global.GVA_DB_LOGIC.Where("id = ?", id).First(&adRewards).Error
	return
}

// @title    GetAdRewardsInfoList
// @description   get AdRewards list by pagination, 分页获取AdRewards
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAdRewardsInfoList(info request.AdRewardsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB_LOGIC.Model(&model.AdRewards{})
	var adRewardss []model.AdRewards
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserId != 0 {
		db = db.Where("`user_id` = ?", info.UserId)
	}
	if info.GuestId != "" {
		db = db.Where("`guest_id` LIKE ?", "%"+info.GuestId+"%")
	}
	if info.AdId != "" {
		db = db.Where("`ad_id` LIKE ?", "%"+info.AdId+"%")
	}
	if info.AdType != "" {
		db = db.Where("`ad_type` = ?", info.AdType)
	}
	if info.AdChannel != "" {
		db = db.Where("`ad_channel` = ?", info.AdChannel)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&adRewardss).Error
	return err, adRewardss, total
}
