package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"sehha/helper"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidationBindJsonField(err error) interface{} {
	var ve validator.ValidationErrors
	var response helper.Response
	if errors.As(err, &ve) {
		out := make([]helper.AppError, len(ve))
		for i, fe := range ve {
			out[i] = helper.AppError{Code: 402, Field: fe.Field(), MessageData: getErrorMsg(fe)}
			response = helper.ResponseError(getErrorMsg(fe), http.StatusBadRequest)
		}
	}
	return response
}

func ValidateField(data interface{}, field string, context *gin.Context) {
	var response helper.Response
	if data == nil || data == "" {
		response = helper.ResponseError(field+" can't be empty", http.StatusBadRequest)
		context.JSON(http.StatusBadRequest, response)
		context.Abort()
	}

}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required " + fe.Field()
	case "lte":
		return "Should be less than" + fe.Field()
	case "gte":
		return "Should be greater than" + fe.Field()
	case "email":
		return "failed email format"
	case "unique":
		return fe.Field() + "already exist"
	}
	return "Unknown error"
}
