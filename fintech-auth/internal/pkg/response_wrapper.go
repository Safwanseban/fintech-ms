package pkg

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, code int, err error, fields ...interface{}) {
	c.JSON(code, ToWrapperError(code, err, fields))
}

type WrapperError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Fields  interface{} `json:"fields"`
}


func ToWrapperError(code int, err error, obj interface{}) WrapperError {
	wrappedErr := WrapperError{
		Code:    fmt.Sprintf("%v", code),
		Message: err.Error(),
		Fields:  obj,
	}

	return wrappedErr
}
