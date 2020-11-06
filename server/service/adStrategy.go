package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
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
    if info.UserRole != "" {
        db = db.Where("`user_role` = ?",info.UserRole)
    }
    if info.Platform != "" {
        db = db.Where("`platform` = ?",info.Platform)
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
    if info.Operator != "" {
        db = db.Where("`operator` LIKE ?","%"+ info.Operator+"%")
    }
	if len(info.Order) > 0 && len(info.OrderBy) > 0 {
		db = db.Order(fmt.Sprintf("%s %s", utils.Camel2Case(info.Order), info.OrderBy))
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&adStrategys).Error
	return err, adStrategys, total
}