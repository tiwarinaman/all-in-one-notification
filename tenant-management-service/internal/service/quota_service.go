package service

import (
	"errors"
	"go.uber.org/zap"
	"tenant-management-service/internal/model"
	"tenant-management-service/internal/model/dto"
	"tenant-management-service/internal/repository"
	"tenant-management-service/pkg/logger"
	"tenant-management-service/pkg/utils"
)

type QuotaService struct {
	repo *repository.QuotaRepository
}

func NewQuotaService(repo *repository.QuotaRepository) *QuotaService {
	return &QuotaService{repo: repo}
}

// UpdateQuotas updates the quotas for a tenant.
func (s *QuotaService) UpdateQuotas(tenantID string, quotas []dto.QuotaDTO) error {
	// Validation
	if err := utils.ValidateNonEmptyString(tenantID, "TenantID"); err != nil {
		return err
	}

	// Convert DTO to model
	var quotaModels []model.Quota
	for _, quota := range quotas {
		quotaModels = append(quotaModels, model.Quota{
			TenantID:     tenantID,
			Channel:      quota.Channel,
			DailyLimit:   quota.DailyLimit,
			MonthlyLimit: quota.MonthlyLimit,
			IsGlobal:     quota.IsGlobal,
		})
	}

	// Call repository to upsert quotas
	if err := s.repo.Upsert(quotaModels); err != nil {
		logger.Error("Error updating quotas", zap.Error(err))
		return errors.New("failed to update quotas")
	}

	return nil
}

// GetQuotas retrieves the quotas for a tenant.
func (s *QuotaService) GetQuotas(tenantID string) ([]model.Quota, error) {
	// Validation
	if err := utils.ValidateNonEmptyString(tenantID, "TenantID"); err != nil {
		return nil, err
	}

	// Fetch quotas from repository
	quotas, err := s.repo.FindByTenantID(tenantID)
	if err != nil {
		logger.Error("Error fetching quotas", zap.Error(err))
		return nil, errors.New("failed to fetch quotas")
	}

	return quotas, nil
}
