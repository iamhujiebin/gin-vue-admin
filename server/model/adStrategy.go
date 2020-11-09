// 自动生成模板AdStrategy
package model

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type AdStrategy struct {
	ID         uint      `gorm:"primarykey"`
	Platform   string    `json:"platform" form:"platform" gorm:"column:platform;comment:平台;type:varchar(32);size:32;"`
	AdPlace    string    `json:"adPlace" form:"adPlace" gorm:"column:ad_place;comment:用于标志具体的页面位置: home:首页 me:我的页面;type:varchar(32);size:32;"`
	AdType     string    `json:"adType" form:"adType" gorm:"column:ad_type;comment:请求的广告的类型 banner：横幅  interstitial：插播广告;type:varchar(32);size:32;"`
	AdPlatform string    `json:"adPlatform" form:"adPlatform" gorm:"column:ad_platform;comment:广告平台，admob、AudienceNetwork等;type:varchar(64);size:64;"`
	AdId       string    `json:"adId" form:"adId" gorm:"column:ad_id;comment:广告id;type:varchar(256);size:256;"`
	Strategy   string    `json:"strategy" form:"strategy" gorm:"column:strategy;comment:广告策略"`
	Enable     *bool     `json:"enable" form:"enable" gorm:"column:enable;comment:是否生效 0：否 1：是;type:tinyint;"`
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime;"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;type:datetime;"`
	Operator   string    `json:"operator" form:"operator" gorm:"column:operator;comment:操作人;type:varchar(256);size:256;"`
}

func (AdStrategy) TableName() string {
	return "ad_strategy"
}
