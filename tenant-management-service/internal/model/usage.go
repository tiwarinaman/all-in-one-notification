package model

import (
	"time"
)

type Usage struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	TenantID          string    `gorm:"size:255;not null" json:"tenant_id"`
	Date              time.Time `gorm:"not null" json:"date"`
	Channel           string    `gorm:"size:50;not null" json:"channel"`
	NotificationsSent int       `gorm:"default:0" json:"notifications_sent"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
