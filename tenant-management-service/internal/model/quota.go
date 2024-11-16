package model

import "time"

type Quota struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	TenantID     string    `gorm:"size:255;not null" json:"tenant_id"`
	Channel      string    `gorm:"size:50;not null" json:"channel"`
	DailyLimit   int       `gorm:"default:10000" json:"daily_limit"`
	MonthlyLimit int       `gorm:"default:300000" json:"monthly_limit"`
	IsGlobal     bool      `gorm:"default:false" json:"is_global"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
