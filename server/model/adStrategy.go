// 自动生成模板AdStrategy
package model

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type AdStrategy struct {
	//global.GVA_MODEL
	ID         uint      `gorm:"primarykey"`
	UserRole   string    `json:"userRole" form:"userRole" gorm:"column:user_role;comment:用户类型,vip/normal等;type:varchar(32);size:32;"`
	Platform   string    `json:"platform" form:"platform" gorm:"column:platform;comment:平台;type:varchar(32);size:32;"`
	Status     *bool     `json:"status" form:"status" gorm:"column:status;comment:是否生效 0：否 1：是;type:tinyint;"`
	Strategy   string    `json:"strategy" form:"strategy" gorm:"column:strategy;comment:广告策略"`
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime;"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;type:datetime;"`
	Operator   string    `json:"operator" form:"operator" gorm:"column:operator;comment:操作人;type:varchar(256);size:256;"`
}

func (AdStrategy) TableName() string {
	return "ad_strategy"
}
