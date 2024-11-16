package repository

import (
	"gorm.io/gorm"
	"tenant-management-service/internal/model"
)

type ConfigRepository struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) *ConfigRepository {
	return &ConfigRepository{db: db}
}

// Upsert inserts or updates configurations in the database.
func (r *ConfigRepository) Upsert(configs []model.Configuration) error {
	for _, config := range configs {
		// Use GORM's Save method to upsert (create or update)
		if err := r.db.Save(&config).Error; err != nil {
			return err
		}
	}
	return nil
}

// FindByTenantId retrieves all configurations for a specific tenant.
func (r *ConfigRepository) FindByTenantId(tenantID string) ([]model.Configuration, error) {
	var configs []model.Configuration
	if err := r.db.Where("tenant_id = ?", tenantID).Find(&configs).Error; err != nil {
		return nil, err
	}
	return configs, nil
}
