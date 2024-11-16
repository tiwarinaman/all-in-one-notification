package model

import "time"

type Tenant struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"size:255;not null" json:"name"`
	ClientID        string    `gorm:"size:255;unique;not null" json:"client_id"`
	ClientSecret    string    `gorm:"size:255;not null" json:"client_secret"`
	Email           string    `gorm:"size:255;not null" json:"email"`
	Phone           string    `gorm:"size:20" json:"phone"`
	Status          string    `gorm:"size:50;default:active" json:"status"`
	BillingTier     string    `gorm:"size:50;default:basic" json:"billing_tier"`
	DefaultLanguage string    `gorm:"size:10;default:'en'" json:"default_language"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
