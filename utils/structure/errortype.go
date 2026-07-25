package structure

import "github.com/labstack/echo/v4"

type ApiResponse struct {
	Response any `json:"response"`
	Code     int
}

func Resp200WithMessage(e echo.Context, msg any) {
	e.JSON(200, ApiResponse{
		Code:     200,
		Response: msg,
	})
}
func Resp400WithMessage(e echo.Context, msg error) {
	e.JSON(400, ApiResponse{
		Code:     400,
		Response: msg.Error(),
	})
}
func Resp500WithMessage(e echo.Context, msg error) {
	e.JSON(500, ApiResponse{
		Code:     500,
		Response: msg.Error(),
	})
}
