package model

import "time"

/**
表模型
 */
type CloserUser struct {
	UserId          string    `gorm:"primary_key" json:"userId"`
	Avatar          string    `json:"avatar"`
	NickName        string    `json:"nickName"`
	Sex             int       `json:"sex"`
	Birthday        time.Time `gorm:"default:null" json:"birthday"`
	Platform        int       `json:"platform"`
	OpenId          string    `json:"openId"`
	UnionId         string    `json:"-"`
	DeviceId        string    `json:"deviceId"`
	Client          string    `json:"user"`
	AppVersion      string    `json:"appVersion"`
	LastOnlineTime  time.Time `gorm:"default:null"`
	LastOfflineTime time.Time `gorm:"default:null"`
	CreateTime      time.Time `gorm:"default:null"`
	IsDelete        int       `gorm:"default:0" json:"-"`
	ForceUpdate     int       `json:"-"`
	PushToken       string    `gorm:"default:''" json:"-"`
	Phone           string    `gorm:"default:''" json:"phone"`
	IsNunar         bool      `gorm:"default:0" json:"is_nunar"`
}