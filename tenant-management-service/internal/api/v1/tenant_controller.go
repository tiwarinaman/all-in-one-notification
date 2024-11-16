package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"tenant-management-service/internal/response"
	"tenant-management-service/internal/service"
	"tenant-management-service/pkg/logger"
)

type TenantController struct {
	service *service.TenantService
}

func NewTenantController(service *service.TenantService) *TenantController {
	return &TenantController{service: service}
}

// Create handles the creation of a new tenant.
func (c *TenantController) Create(ctx *gin.Context) {
	var req struct {
		Name            string `json:"name" binding:"required"`
		Email           string `json:"email" binding:"required,email"`
		Phone           string `json:"phone" binding:"required"`
		BillingTier     string `json:"billing_tier" binding:"required"`
		DefaultLanguage string `json:"default_language" binding:"required"`
	}

	// Bind the JSON request body to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid input in CreateTenant", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Invalid input", "INVALID_INPUT", err.Error())
		return
	}

	// Call the service to create the tenant
	tenant, err := c.service.CreateTenant(req.Name, req.Email, req.Phone, req.BillingTier, req.DefaultLanguage)
	if err != nil {
		logger.Error("Failed to create tenant", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Failed to create tenant", "CREATE_FAILED", err.Error())
		return
	}

	logger.Info("Tenant created successfully", zap.Uint("tenant_id", tenant.ID))
	response.Success(ctx, 201, "Tenant created successfully", tenant, nil)
}

// Get handles fetching details of a tenant by ID.
func (c *TenantController) Get(ctx *gin.Context) {
	// Parse the tenant ID from the URL
	id, err := strconv.Atoi(ctx.Param("tenant_id"))
	if err != nil {
		logger.Warn("Invalid tenant ID in GetTenant", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Invalid tenant ID", "INVALID_ID", err.Error())
		return
	}

	// Call the service to fetch the tenant details
	tenant, err := c.service.GetTenantByID(uint(id))
	if err != nil {
		logger.Error("Tenant not found", zap.Int("tenant_id", id), zap.Error(err))
		response.Error(ctx, http.StatusNotFound, "Tenant not found", "NOT_FOUND", err.Error())
		return
	}

	logger.Info("Tenant retrieved successfully", zap.Int("tenant_id", id))
	response.Success(ctx, 200, "Tenant retrieved successfully", tenant, nil)
}

// Update handles updating an existing tenant's details.
func (c *TenantController) Update(ctx *gin.Context) {
	// Parse the tenant ID from the URL
	id, err := strconv.Atoi(ctx.Param("tenant_id"))
	if err != nil {
		logger.Warn("Invalid tenant ID in UpdateTenant", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Invalid tenant ID", "INVALID_ID", err.Error())
		return
	}

	var req struct {
		Name            string `json:"name"`
		Email           string `json:"email" binding:"email"`
		Phone           string `json:"phone"`
		BillingTier     string `json:"billing_tier"`
		DefaultLanguage string `json:"default_language"`
	}

	// Bind the JSON request body to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid input in UpdateTenant", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Invalid input", "INVALID_INPUT", err.Error())
		return
	}

	// Call the service to update the tenant
	err = c.service.UpdateTenant(uint(id), req.Name, req.Email, req.Phone, req.BillingTier, req.DefaultLanguage)
	if err != nil {
		logger.Error("Failed to update tenant", zap.Int("tenant_id", id), zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Failed to update tenant", "UPDATE_FAILED", err.Error())
		return
	}

	logger.Info("Tenant updated successfully", zap.Int("tenant_id", id))
	response.Success(ctx, 200, "Tenant updated successfully", nil, nil)
}

// Delete handles deleting a tenant by ID.
func (c *TenantController) Delete(ctx *gin.Context) {
	// Parse the tenant ID from the URL
	id, err := strconv.Atoi(ctx.Param("tenant_id"))
	if err != nil {
		logger.Warn("Invalid tenant ID in DeleteTenant", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Invalid tenant ID", "INVALID_ID", err.Error())
		return
	}

	// Call the service to delete the tenant
	err = c.service.DeleteTenant(uint(id))
	if err != nil {
		logger.Error("Failed to delete tenant", zap.Int("tenant_id", id), zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Failed to delete tenant", "DELETE_FAILED", err.Error())
		return
	}

	logger.Info("Tenant deleted successfully", zap.Int("tenant_id", id))
	response.Success(ctx, 204, "Tenant deleted successfully", nil, nil)
}
