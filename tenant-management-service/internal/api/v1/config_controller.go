package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tenant-management-service/internal/response"
	"tenant-management-service/internal/service"
	"tenant-management-service/pkg/logger"
)

type ConfigController struct {
	service *service.ConfigService
}

func NewConfigController(service *service.ConfigService) *ConfigController {
	return &ConfigController{service: service}
}

// UpsertConfig handles creating or updating configurations for a tenant.
func (c *ConfigController) UpsertConfig(ctx *gin.Context) {

	tenantId := ctx.Param("tenant_id")
	var configs []struct {
		ConfigKey   string `json:"config_key" binding:"required"`
		ConfigValue string `json:"config_value" binding:"required"`
		IsGlobal    bool   `json:"is_global"`
	}

	// Validate input
	if err := ctx.ShouldBindJSON(&configs); err != nil {
		logger.Warn("Invalid input in UpsertConfig", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Invalid input", "INVALID_INPUT", err.Error())
		return
	}

	// Call service to upsert configurations
	err := c.service.UpsertConfigurations(tenantId, []struct {
		ConfigKey   string
		ConfigValue string
		IsGlobal    bool
	}(configs))
	if err != nil {
		logger.Error("Failed to upsert configurations", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Failed to upsert configurations", "UPSERT_FAILED", err.Error())
		return
	}

	logger.Info("Configurations upserted successfully", zap.String("tenant_id", tenantId))
	response.Success(ctx, 200, "Configurations upserted successfully", nil, nil)
}

// GetConfigs retrieves all configurations for a tenant.
func (c *ConfigController) GetConfigs(ctx *gin.Context) {
	tenantID := ctx.Param("tenant_id")

	// Call service to retrieve configurations
	configs, err := c.service.GetConfigurations(tenantID)
	if err != nil {
		logger.Error("Failed to fetch configurations", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Failed to fetch configurations", "FETCH_FAILED", err.Error())
		return
	}

	logger.Info("Configurations retrieved successfully", zap.String("tenant_id", tenantID))
	response.Success(ctx, 200, "Configurations retrieved successfully", configs, nil)
}
