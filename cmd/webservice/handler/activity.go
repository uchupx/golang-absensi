package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/uchupx/golang-absensi/internal/constant"
	"github.com/uchupx/golang-absensi/pkg/dto"
	"github.com/uchupx/golang-absensi/pkg/errors"
	"github.com/uchupx/golang-absensi/pkg/util/httputil"
	"github.com/uchupx/golang-absensi/pkg/util/textutil"
)

type FindActivityByUserId func(ctx context.Context, req dto.ActivityGetRequest) (resp dto.BaseListResponse, err error)
type CreateActivity func(ctx context.Context, payload dto.ActivityPostPayload) (resp *string, err error)
type EditActivity func(ctx context.Context, id uint64, payload dto.ActivityPutPayload) (resp *string, err error)
type DeleteActivity func(ctx context.Context, id uint64) (resp *string, err error)

func HandlerGetActvity(handler FindActivityByUserId) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request dto.ActivityGetRequest

		if err := c.Bind(&request); err != nil {
			return httputil.WriteErrorResponse(c, errors.ErrUnparseableRequestBody)
		}

		user := c.Get(constant.UserKeyCtx)
		ctx := context.WithValue(c.Request().Context(), constant.UserKeyCtx, user)

		resp, err := handler(ctx, request)
		if err != nil {
			return httputil.WriteErrorResponse(c, err)
		}

		return httputil.WriteSuccessResponse(c, resp)
	}
}

func HandlerCreateActivity(handler CreateActivity) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request dto.ActivityPostPayload

		err := c.Bind(&request)
		if err != nil {
			return httputil.WriteErrorResponse(c, errors.ErrUnparseableRequestBody)
		}

		user := c.Get(constant.UserKeyCtx)
		ctx := context.WithValue(c.Request().Context(), constant.UserKeyCtx, user)

		resp, err := handler(ctx, request)
		if err != nil {
			return httputil.WriteErrorResponse(c, err)
		}

		return httputil.WriteSuccessResponse(c, resp)
	}
}

func HandlerUpdateActivity(handler EditActivity) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request dto.ActivityPutPayload
		idStr := c.Param("id")

		err := c.Bind(&request)
		if err != nil {
			return httputil.WriteErrorResponse(c, errors.ErrUnparseableRequestBody)
		}

		id := textutil.StringToUint64(idStr, 0)

		user := c.Get(constant.UserKeyCtx)
		ctx := context.WithValue(c.Request().Context(), constant.UserKeyCtx, user)

		resp, err := handler(ctx, id, request)
		if err != nil {
			return httputil.WriteErrorResponse(c, err)
		}

		return httputil.WriteSuccessResponse(c, resp)
	}
}

func HandlerDeleteActivity(handler DeleteActivity) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := textutil.StringToUint64(c.Param("id"), 0)

		user := c.Get(constant.UserKeyCtx)
		ctx := context.WithValue(c.Request().Context(), constant.UserKeyCtx, user)

		resp, err := handler(ctx, id)
		if err != nil {
			return httputil.WriteErrorResponse(c, err)
		}

		return httputil.WriteSuccessResponse(c, resp)
	}
}
