package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/uchupx/golang-absensi/internal/constant"
	"github.com/uchupx/golang-absensi/pkg/dto"
	"github.com/uchupx/golang-absensi/pkg/util/httputil"
)

type FindAttendance func(ctx context.Context) (resp dto.BaseListResponse, err error)
type CheckIn func(ctx context.Context) (resp *string, err error)
type CheckOut func(ctx context.Context) (resp *string, err error)

func HandlerFindAttendance(handler FindAttendance) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get(constant.UserKeyCtx)
		ctx := context.WithValue(c.Request().Context(), constant.UserKeyCtx, user)

		resp, err := handler(ctx)
		if err != nil {
			return httputil.WriteErrorResponse(c, err)
		}

		return httputil.WriteSuccessResponse(c, resp)
	}
}

func HandlerCheckIn(handler CheckIn) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get(constant.UserKeyCtx)
		ctx := context.WithValue(c.Request().Context(), constant.UserKeyCtx, user)

		resp, err := handler(ctx)
		if err != nil {
			return httputil.WriteErrorResponse(c, err)
		}

		return httputil.WriteSuccessResponse(c, resp)
	}
}

func HandlerCheckOut(handler CheckOut) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get(constant.UserKeyCtx)
		ctx := context.WithValue(c.Request().Context(), constant.UserKeyCtx, user)

		resp, err := handler(ctx)
		if err != nil {
			return httputil.WriteErrorResponse(c, err)
		}

		return httputil.WriteSuccessResponse(c, resp)
	}
}
