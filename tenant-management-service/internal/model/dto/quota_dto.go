package dto

type QuotaDTO struct {
	Channel      string `json:"channel" binding:"required"`
	DailyLimit   int    `json:"daily_limit" binding:"required"`
	MonthlyLimit int    `json:"monthly_limit" binding:"required"`
	IsGlobal     bool   `json:"is_global"`
}
