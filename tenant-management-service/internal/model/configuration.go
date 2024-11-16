package model

import "time"

type Configuration struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TenantID    string    `gorm:"size:255;not null" json:"tenant_id"`
	ConfigKey   string    `gorm:"size:255;not null" json:"config_key"`
	ConfigValue string    `gorm:"size:255;not null" json:"config_value"`
	IsGlobal    bool      `gorm:"default:false" json:"is_global"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
