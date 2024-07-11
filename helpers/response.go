package helpers

import (
	"github.com/arioprima/jobseekers_api/schemas"
	"github.com/gin-gonic/gin"
)

func ApiResponse(ctx *gin.Context, Code int, Status string, Message string, Data interface{}, Auth interface{}) {
	jsonResponse := schemas.SchemaResponses{
		Code:    Code,
		Status:  Status,
		Message: Message,
		Data:    Data,
		Auth:    Auth,
	}

	if Code >= 400 {
		ctx.AbortWithStatusJSON(Code, jsonResponse)
	} else {
		ctx.JSON(Code, jsonResponse)
	}
}

func ValidatorErrorResponse(ctx *gin.Context, Code int, Status string, Error interface{}) {
	errResponse := schemas.SchemaErrorResponse{
		Code:   Code,
		Status: Status,
		Error:  Error,
	}

	ctx.AbortWithStatusJSON(Code, errResponse)
}
