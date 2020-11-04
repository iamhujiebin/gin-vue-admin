// 自动生成模板AdReports
package model

import "gin-vue-admin/global"

// 如果含有time.Time 请自行import time包
type AdReports struct {
	global.GVA_MODEL
	Platform string `json:"platform" form:"platform" gorm:"column:platform;comment:平台，google/apple;type:varchar(45);size:45;"`
	AdId     string `json:"adId" form:"adId" gorm:"column:ad_id;comment:广告id;type:varchar(256);size:256;"`
	AdType   string `json:"adType" form:"adType" gorm:"column:ad_type;comment:广告类型;type:varchar(45);size:45;"`
	UserId   int    `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;type:int;size:10;"`
	GuestId  string `json:"guestId" form:"guestId" gorm:"column:guest_id;comment:设备id;type:varchar(256);size:256;"`
}

func (AdReports) TableName() string {
	return "ad_reports"
}
