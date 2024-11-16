package response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(ctx *gin.Context, statusCode int, message string, data, meta interface{}) {
	ctx.JSON(statusCode, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func Error(ctx *gin.Context, statusCode int, message string, errorCode string, details interface{}) {
	ctx.JSON(statusCode, APIResponse{
		Status:  "error",
		Message: message,
		Error: gin.H{
			"code":    errorCode,
			"details": details,
		},
	})
}
