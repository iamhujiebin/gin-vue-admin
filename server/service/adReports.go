package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateAdReports
// @description   create a AdReports
// @param     adReports               model.AdReports
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAdReports(adReports model.AdReports) (err error) {
	err = global.GVA_DB_LOGIC.Create(&adReports).Error
	return err
}

// @title    DeleteAdReports
// @description   delete a AdReports
// @auth                     （2020/04/05  20:22）
// @param     adReports               model.AdReports
// @return                    error

func DeleteAdReports(adReports model.AdReports) (err error) {
	err = global.GVA_DB_LOGIC.Delete(adReports).Error
	return err
}

// @title    DeleteAdReportsByIds
// @description   delete AdReportss
// @auth                     （2020/04/05  20:22）
// @param     adReports               model.AdReports
// @return                    error

func DeleteAdReportsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB_LOGIC.Delete(&[]model.AdReports{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateAdReports
// @description   update a AdReports
// @param     adReports          *model.AdReports
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAdReports(adReports *model.AdReports) (err error) {
	err = global.GVA_DB_LOGIC.Save(adReports).Error
	return err
}

// @title    GetAdReports
// @description   get the info of a AdReports
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    AdReports        AdReports

func GetAdReports(id uint) (err error, adReports model.AdReports) {
	err = global.GVA_DB_LOGIC.Where("id = ?", id).First(&adReports).Error
	return
}

// @title    GetAdReportsInfoList
// @description   get AdReports list by pagination, 分页获取AdReports
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAdReportsInfoList(info request.AdReportsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB_LOGIC.Model(&model.AdReports{})
	var adReportss []model.AdReports
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
	if info.UserId != 0 {
		db = db.Where("`user_id` = ?", info.UserId)
	}
	if info.GuestId != "" {
		db = db.Where("`guest_id` LIKE ?", "%"+info.GuestId+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&adReportss).Error
	return err, adReportss, total
}
