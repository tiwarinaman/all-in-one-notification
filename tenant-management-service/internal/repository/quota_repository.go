package repository

import (
	"gorm.io/gorm"
	"tenant-management-service/internal/model"
)

type QuotaRepository struct {
	db *gorm.DB
}

func NewQuotaRepository(db *gorm.DB) *QuotaRepository {
	return &QuotaRepository{db: db}
}

// Upsert inserts or updates quotas in the database.
func (r *QuotaRepository) Upsert(quotas []model.Quota) error {
	for _, quota := range quotas {
		// Use GORM's Save method to upsert (create or update)
		if err := r.db.Save(&quota).Error; err != nil {
			return err
		}
	}
	return nil
}

// FindByTenantID retrieves all quotas for a specific tenant.
func (r *QuotaRepository) FindByTenantID(tenantID string) ([]model.Quota, error) {
	var quotas []model.Quota
	if err := r.db.Where("tenant_id = ?", tenantID).Find(&quotas).Error; err != nil {
		return nil, err
	}
	return quotas, nil
}
