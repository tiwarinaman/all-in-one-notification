package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tenant-management-service/internal/response"
	"tenant-management-service/internal/service"
	"tenant-management-service/pkg/logger"
)

type UsageController struct {
	service *service.UsageService
}

func NewUsageController(service *service.UsageService) *UsageController {
	return &UsageController{service: service}
}

// GetUsage retrieves usage data for a tenant.
func (c *UsageController) GetUsage(ctx *gin.Context) {
	tenantID := ctx.Param("tenant_id")
	channel := ctx.Query("channel") // Optional query parameter to filter by channel

	// Fetch usage data
	usage, err := c.service.GetUsage(tenantID, channel)
	if err != nil {
		logger.Error("Failed to fetch usage data", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Failed to fetch usage data", "FETCH_FAILED", err.Error())
		return
	}

	logger.Info("Usage data retrieved successfully", zap.String("tenant_id", tenantID), zap.String("channel", channel))
	response.Success(ctx, http.StatusOK, "Usage data retrieved successfully", usage, nil)
}
