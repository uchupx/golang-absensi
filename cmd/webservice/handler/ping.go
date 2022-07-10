package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/uchupx/golang-absensi/pkg/dto"
	"github.com/uchupx/golang-absensi/pkg/util/httputil"
)

type Ping func(ctx context.Context) (resp dto.BaseResponse, err error)

func HandlerPing(handler Ping) echo.HandlerFunc {
	return func(c echo.Context) error {
		resp, err := handler(c.Request().Context())
		if err != nil {
			panic(err)
		}

		return httputil.WriteSuccessResponse(c, resp)
	}
}
