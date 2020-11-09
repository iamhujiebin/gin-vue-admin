package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
)

// @title    CreateAdEvents
// @description   create a AdEvents
// @param     adEvents               model.AdEvents
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAdEvents(adEvents model.AdEvents) (err error) {
	err = global.GVA_DB_LOGIC.Create(&adEvents).Error
	return err
}

// @title    DeleteAdEvents
// @description   delete a AdEvents
// @auth                     （2020/04/05  20:22）
// @param     adEvents               model.AdEvents
// @return                    error

func DeleteAdEvents(adEvents model.AdEvents) (err error) {
	err = global.GVA_DB_LOGIC.Delete(adEvents).Error
	return err
}

// @title    DeleteAdEventsByIds
// @description   delete AdEventss
// @auth                     （2020/04/05  20:22）
// @param     adEvents               model.AdEvents
// @return                    error

func DeleteAdEventsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB_LOGIC.Delete(&[]model.AdEvents{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateAdEvents
// @description   update a AdEvents
// @param     adEvents          *model.AdEvents
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAdEvents(adEvents *model.AdEvents) (err error) {
	err = global.GVA_DB_LOGIC.Save(adEvents).Error
	return err
}

// @title    GetAdEvents
// @description   get the info of a AdEvents
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    AdEvents        AdEvents

func GetAdEvents(id uint) (err error, adEvents model.AdEvents) {
	err = global.GVA_DB_LOGIC.Where("id = ?", id).First(&adEvents).Error
	return
}

// @title    GetAdEventsInfoList
// @description   get AdEvents list by pagination, 分页获取AdEvents
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAdEventsInfoList(info request.AdEventsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB_LOGIC.Model(&model.AdEvents{})
	var adEventss []model.AdEvents
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Platform != "" {
		db = db.Where("`platform` = ?", info.Platform)
	}
	if info.AdId != "" {
		db = db.Where("`ad_id` LIKE ?", "%"+info.AdId+"%")
	}
	if info.AdType != "" {
		db = db.Where("`ad_type` = ?", info.AdType)
	}
	if info.AdAction != "" {
		db = db.Where("`ad_action` = ?", info.AdAction)
	}
	if info.AdChannel != "" {
		db = db.Where("`ad_channel` = ?", info.AdChannel)
	}
	if !info.CreateTime.IsZero() {
		db = db.Where("`create_time` <> ?", info.CreateTime)
	}
	if !info.UpdateTime.IsZero() {
		db = db.Where("`update_time` <> ?", info.UpdateTime)
	}
	if len(info.Order) > 0 && len(info.OrderBy) > 0 {
		db = db.Order(fmt.Sprintf("%s %s", utils.Camel2Case(info.Order), info.OrderBy))
	} else {
		db = db.Order("id desc")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&adEventss).Error
	return err, adEventss, total
}
