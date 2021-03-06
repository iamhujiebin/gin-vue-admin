package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateAdStrategy
// @description   create a AdStrategy
// @param     adStrategy               model.AdStrategy
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAdStrategy(adStrategy model.AdStrategy) (err error) {
	err = global.GVA_DB_LOGIC.Create(&adStrategy).Error
	return err
}

// @title    DeleteAdStrategy
// @description   delete a AdStrategy
// @auth                     （2020/04/05  20:22）
// @param     adStrategy               model.AdStrategy
// @return                    error

func DeleteAdStrategy(adStrategy model.AdStrategy) (err error) {
	err = global.GVA_DB_LOGIC.Delete(adStrategy).Error
	return err
}

// @title    DeleteAdStrategyByIds
// @description   delete AdStrategys
// @auth                     （2020/04/05  20:22）
// @param     adStrategy               model.AdStrategy
// @return                    error

func DeleteAdStrategyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB_LOGIC.Delete(&[]model.AdStrategy{},"id in ?",ids.Ids).Error
	return err
}

// @title    UpdateAdStrategy
// @description   update a AdStrategy
// @param     adStrategy          *model.AdStrategy
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAdStrategy(adStrategy *model.AdStrategy) (err error) {
	err = global.GVA_DB_LOGIC.Save(adStrategy).Error
	return err
}

// @title    GetAdStrategy
// @description   get the info of a AdStrategy
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    AdStrategy        AdStrategy

func GetAdStrategy(id uint) (err error, adStrategy model.AdStrategy) {
	err = global.GVA_DB_LOGIC.Where("id = ?", id).First(&adStrategy).Error
	return
}

// @title    GetAdStrategyInfoList
// @description   get AdStrategy list by pagination, 分页获取AdStrategy
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAdStrategyInfoList(info request.AdStrategySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_LOGIC.Model(&model.AdStrategy{})
    var adStrategys []model.AdStrategy
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Platform != "" {
        db = db.Where("`platform` LIKE ?","%"+ info.Platform+"%")
    }
    if info.AdPlace != "" {
        db = db.Where("`ad_place` = ?",info.AdPlace)
    }
    if info.AdType != "" {
        db = db.Where("`ad_type` = ?",info.AdType)
    }
    if info.AdPlatform != "" {
        db = db.Where("`ad_platform` LIKE ?","%"+ info.AdPlatform+"%")
    }
    if info.AdId != "" {
        db = db.Where("`ad_id` LIKE ?","%"+ info.AdId+"%")
    }
    if info.Enable != nil {
        db = db.Where("`enable` = ?",info.Enable)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&adStrategys).Error
	return err, adStrategys, total
}