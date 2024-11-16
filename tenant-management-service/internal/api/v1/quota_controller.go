package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tenant-management-service/internal/model/dto"
	"tenant-management-service/internal/response"
	"tenant-management-service/internal/service"
	"tenant-management-service/pkg/logger"
)

type QuotaController struct {
	service *service.QuotaService
}

func NewQuotaController(service *service.QuotaService) *QuotaController {
	return &QuotaController{service: service}
}

// UpdateQuota handles updating quotas for a tenant.
func (c *QuotaController) UpdateQuota(ctx *gin.Context) {
	tenantID := ctx.Param("tenant_id")
	var quotas []dto.QuotaDTO

	// Validate input
	if err := ctx.ShouldBindJSON(&quotas); err != nil {
		logger.Warn("Invalid input in UpdateQuota", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Invalid input", "INVALID_INPUT", err.Error())
		return
	}

	// Call service to update quotas
	err := c.service.UpdateQuotas(tenantID, quotas)
	if err != nil {
		logger.Error("Failed to update quotas", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Failed to update quotas", "UPDATE_FAILED", err.Error())
		return
	}

	logger.Info("Quotas updated successfully", zap.String("tenant_id", tenantID))
	response.Success(ctx, http.StatusOK, "Quotas updated successfully", nil, nil)
}

// GetQuotas retrieves all quotas for a tenant.
func (c *QuotaController) GetQuotas(ctx *gin.Context) {
	tenantID := ctx.Param("tenant_id")

	// Call service to fetch quotas
	quotas, err := c.service.GetQuotas(tenantID)
	if err != nil {
		logger.Error("Failed to fetch quotas", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Failed to fetch quotas", "FETCH_FAILED", err.Error())
		return
	}

	logger.Info("Quotas retrieved successfully", zap.String("tenant_id", tenantID))
	response.Success(ctx, http.StatusOK, "Quotas retrieved successfully", quotas, nil)
}
