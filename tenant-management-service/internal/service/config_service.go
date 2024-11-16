package service

import (
	"errors"
	"go.uber.org/zap"
	"tenant-management-service/internal/model"
	"tenant-management-service/internal/repository"
	"tenant-management-service/pkg/logger"
	"tenant-management-service/pkg/utils"
)

type ConfigService struct {
	repo *repository.ConfigRepository
}

func NewConfigService(repo *repository.ConfigRepository) *ConfigService {
	return &ConfigService{repo: repo}
}

// UpsertConfigurations creates or updates configurations for a tenant.
func (s *ConfigService) UpsertConfigurations(tenantID string, configs []struct {
	ConfigKey   string
	ConfigValue string
	IsGlobal    bool
}) error {

	// Validation
	if err := utils.ValidateNonEmptyString(tenantID, "tenantId"); err != nil {
		return err
	}

	// Convert input to model
	var configModels []model.Configuration
	for _, config := range configs {
		configModels = append(configModels, model.Configuration{
			TenantID:    tenantID,
			ConfigKey:   config.ConfigKey,
			ConfigValue: config.ConfigValue,
			IsGlobal:    config.IsGlobal,
		})
	}

	// Call repository to upsert configurations
	if err := s.repo.Upsert(configModels); err != nil {
		logger.Error("Error upserting configurations", zap.Error(err))
		return errors.New("failed to upsert configurations")
	}

	return nil
}

// GetConfigurations retrieves configurations for a tenant.
func (s *ConfigService) GetConfigurations(tenantId string) ([]model.Configuration, error) {

	// Validation
	if err := utils.ValidateNonEmptyString(tenantId, "tenantId"); err != nil {
		return nil, err
	}

	// Fetch configurations from repository
	configs, err := s.repo.FindByTenantId(tenantId)
	if err != nil {
		logger.Error("Error fetching configurations", zap.Error(err))
		return nil, errors.New("failed to fetch configurations")
	}

	return configs, nil
}
