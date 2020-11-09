// 自动生成模板AdEvents
package model

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type AdEvents struct {
	//global.GVA_MODEL
	ID         uint      `gorm:"primarykey"`
	Platform   string    `json:"platform" form:"platform" gorm:"column:platform;comment:平台，google/ios;type:varchar(45);size:45;"`
	AdId       string    `json:"adId" form:"adId" gorm:"column:ad_id;comment:广告id;type:varchar(256);size:256;"`
	AdType     string    `json:"adType" form:"adType" gorm:"column:ad_type;comment:广告类型，video/banner/gift等;type:varchar(45);size:45;"`
	AdAction   string    `json:"adAction" form:"adAction" gorm:"column:ad_action;comment:广告动作，如加载、点击、关闭等;type:varchar(256);size:256;"`
	AdPlatform string    `json:"adPlatform" form:"adPlatform" gorm:"column:ad_platform;comment:广告平台，admob、AudienceNetwork等;type:varchar(256);size:256;"`
	UserId     int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;type:int;size:10;"`
	GuestId    string    `json:"guestId" form:"guestId" gorm:"column:guest_id;comment:设备id;type:varchar(256);size:256;"`
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;type:datetime;"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;type:datetime;"`
}

func (AdEvents) TableName() string {
	return "ad_events"
}
