package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type PaginationResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	Total   int64       `json:"total"`
}

func Success(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Status:  "failed",
		Code:    code,
		Data:    nil,
		Message: message,
	})
}

func SuccessPagination(c *gin.Context, data interface{}, code, page, size int, total int64) {
	c.JSON(http.StatusOK, PaginationResponse{
		Status: "success",
		Code:   code,
		Data:   data,
		Page:   page,
		Size:   size,
		Total:  total,
	})
}
