package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tenant-management-service/internal/response"
	"tenant-management-service/internal/service"
	pkgerr "tenant-management-service/pkg/error"
	"tenant-management-service/pkg/logger"
)

// AuthMiddleware validates client_id and client_secret in the headers.
func AuthMiddleware(tenantService *service.TenantService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		clientId := ctx.GetHeader("X-Client-Id")
		clientSecret := ctx.GetHeader("X-Client-Secret")

		if clientId == "" || clientSecret == "" {
			logger.Warn("Missing X-Client-Id or X-Client-Secret in headers")
			response.Error(
				ctx,
				http.StatusUnauthorized,
				pkgerr.ErrUnauthorized.Error(),
				"UNAUTHORIZED",
				"Missing X-Client-Id or X-Client-Secret in headers",
			)
			ctx.Abort()
			return
		}

		// Validate client_id and client_secret using the tenant service
		isValid, err := tenantService.ValidateClientCredentials(clientId, clientSecret)
		if err != nil {
			logger.Error("Error validating client credentials", zap.Error(err))
			response.Error(
				ctx,
				http.StatusInternalServerError,
				pkgerr.ErrInternalServer.Error(),
				"INTERNAL_SERVER_ERROR",
				nil,
			)
			ctx.Abort()
			return
		}

		if !isValid {
			logger.Warn("Invalid client_id or client_secret", zap.String("client_id", clientId))
			response.Error(
				ctx,
				http.StatusInternalServerError,
				pkgerr.ErrUnauthorized.Error(),
				"UNAUTHORIZED",
				"Invalid client_id or client_secret",
			)
			ctx.Abort()
			return
		}

		// Proceed to the next handler if validation is successful
		ctx.Next()
	}
}
