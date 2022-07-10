package httputil

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uchupx/golang-absensi/pkg/dto"
	"github.com/uchupx/golang-absensi/pkg/errors"
)

func WriteSuccessResponse(e echo.Context, payload interface{}) error {
	return WriteResponse(e, dto.ResponseParam{
		Status: http.StatusOK,
		Payload: dto.BaseResponse{
			Data: payload,
		},
	})
}

func WriteErrorResponse(e echo.Context, er error) error {
	errResp := errors.GetErrorResponse(er)
	return WriteResponse(e, dto.ResponseParam{
		Status: int(errResp.HTTPCode),
		Payload: dto.BaseResponse{
			Error: &dto.ErrorResponse{
				Code:    errResp.Code,
				Message: errResp.Message,
			},
		},
	})
}

func WriteResponse(e echo.Context, param dto.ResponseParam) error {
	return e.JSON(param.Status, param.Payload)
}
