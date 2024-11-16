package service

import (
	"errors"
	"go.uber.org/zap"
	"tenant-management-service/internal/model"
	"tenant-management-service/internal/repository"
	"tenant-management-service/pkg/logger"
	"tenant-management-service/pkg/utils"
)

type UsageService struct {
	repo *repository.UsageRepository
}

func NewUsageService(repo *repository.UsageRepository) *UsageService {
	return &UsageService{repo: repo}
}

// GetUsage retrieves usage data for a tenant, optionally filtered by channel.
func (s *UsageService) GetUsage(tenantID string, channel string) ([]model.Usage, error) {
	// Validation
	if err := utils.ValidateNonEmptyString(tenantID, "TenantID"); err != nil {
		return nil, err
	}

	// Fetch usage data from repository
	usage, err := s.repo.FindByTenantIDAndChannel(tenantID, channel)
	if err != nil {
		logger.Error("Error fetching usage data", zap.Error(err))
		return nil, errors.New("failed to fetch usage data")
	}

	return usage, nil
}
