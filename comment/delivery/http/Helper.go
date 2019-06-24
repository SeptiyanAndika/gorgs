package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ResponseObject struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}


func successResponse(ctx echo.Context, message string, data interface{}) error {
	return ctx.JSON(http.StatusOK, ResponseObject{
		Success: true,
		Message: message,
		Data:    data,
	})

}

func errorResponse(ctx echo.Context, err error, data interface{}) error {
	return ctx.JSON(http.StatusBadRequest, ResponseObject{
		Success: false,
		Message: err.Error(),
		Data:    data,
	})
}
