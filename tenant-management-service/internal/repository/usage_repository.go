package repository

import (
	"gorm.io/gorm"
	"tenant-management-service/internal/model"
)

type UsageRepository struct {
	db *gorm.DB
}

func NewUsageRepository(db *gorm.DB) *UsageRepository {
	return &UsageRepository{db: db}
}

// FindByTenantIDAndChannel retrieves usage data for a tenant, optionally filtered by channel.
func (r *UsageRepository) FindByTenantIDAndChannel(tenantID string, channel string) ([]model.Usage, error) {
	var usage []model.Usage
	query := r.db.Where("tenant_id = ?", tenantID)

	// Filter by channel if provided
	if channel != "" {
		query = query.Where("channel = ?", channel)
	}

	if err := query.Find(&usage).Error; err != nil {
		return nil, err
	}
	return usage, nil
}
