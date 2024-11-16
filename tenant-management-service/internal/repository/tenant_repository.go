package repository

import (
	"errors"
	"gorm.io/gorm"
	"tenant-management-service/internal/model"
	pkgerr "tenant-management-service/pkg/error"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) Create(tenant *model.Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *TenantRepository) FindById(id uint) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.First(&tenant, id).Error
	return &tenant, err
}

func (r *TenantRepository) Update(tenant *model.Tenant) error {
	return r.db.Save(tenant).Error
}

func (r *TenantRepository) Delete(id uint) error {
	return r.db.Delete(&model.Tenant{}, id).Error
}

// FindByClientId retrieves a tenant by its client_id.
func (r *TenantRepository) FindByClientId(clientId string) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.Where("client_id = ?", clientId).First(&tenant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkgerr.ErrNotFound
		}
		return nil, err
	}
	return &tenant, nil
}
