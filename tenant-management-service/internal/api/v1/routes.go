package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tenant-management-service/internal/repository"
	"tenant-management-service/internal/service"
	"tenant-management-service/pkg/middleware"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {

	// Initialize repositories
	tenantRepo := repository.NewTenantRepository(db)
	configRepo := repository.NewConfigRepository(db)
	quotaRepo := repository.NewQuotaRepository(db)
	usageRepo := repository.NewUsageRepository(db)

	// Initialize services
	tenantService := service.NewTenantService(tenantRepo)
	configService := service.NewConfigService(configRepo)
	quotaService := service.NewQuotaService(quotaRepo)
	usageService := service.NewUsageService(usageRepo)

	// Initialize controllers
	tenantController := NewTenantController(tenantService)
	configController := NewConfigController(configService)
	quotaController := NewQuotaController(quotaService)
	usageController := NewUsageController(usageService)

	// Define routes
	api := router.Group("/api/v1")

	// Public Route
	api.POST("/tenants", tenantController.Create)

	// Protected Routes
	protected := api.Use(middleware.AuthMiddleware(tenantService))
	{
		// Tenant Management Routes
		protected.GET("/tenants/:tenant_id", tenantController.Get)
		protected.PUT("/tenants/:tenant_id", tenantController.Update)
		protected.DELETE("/tenants/:tenant_id", tenantController.Delete)

		// Configuration Management Routes
		protected.PUT("/tenants/:tenant_id/configs", configController.UpsertConfig)
		protected.GET("/tenants/:tenant_id/configs", configController.GetConfigs)

		// Quota Management Routes
		protected.PUT("/tenants/:tenant_id/quotas", quotaController.UpdateQuota)
		protected.GET("/tenants/:tenant_id/quotas", quotaController.GetQuotas)

		// Usage Management Routes
		protected.GET("/tenants/:tenant_id/usage", usageController.GetUsage)
	}

}
