package models

import "time"

type ProviderInfo struct {
	ID             uint
	AppProviderID  string `gorm:"size:16;not null"`
	AppAuthCode    string `gorm:"size:32;not null"`
	ProviderName   string `gorm:"size:50"`
	ContactAddress string `gorm:"size:100"`
	Contacts       string `gorm:"size:10"`
	Phone          string `gorm:"size:50"`
	Email          string `gorm:"size:50"`
	IsOpenSDK      uint8  `gorm:"not null"`
	HighDefense    string `gorm:"size:100"`
	AuthorizedTime time.Time
	ExpireTime     time.Time
}

func (ProviderInfo) TableName() string {
	return "provider_info"
}
