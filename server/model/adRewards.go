// 自动生成模板AdRewards
package model

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type AdRewards struct {
	//global.GVA_MODEL
	ID           uint      `gorm:"primarykey"`
	UserId       int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;type:int;size:10;"`
	GuestId      string    `json:"guestId" form:"guestId" gorm:"column:guest_id;comment:设备id;type:varchar(256);size:256;"`
	AdId         string    `json:"adId" form:"adId" gorm:"column:ad_id;comment:广告id;type:varchar(256);size:256;"`
	CreateTime   time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime;"`
	UpdateTime   time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;type:datetime;"`
	RewardDetail string    `json:"rewardDetail" form:"rewardDetail" gorm:"column:reward_detail;comment:奖励明细;type:json;"`
	AdType       string    `json:"adType" form:"adType" gorm:"column:ad_type;comment:广告类型;type:varchar(256);size:256;"`
}

func (AdRewards) TableName() string {
	return "ad_rewards"
}
