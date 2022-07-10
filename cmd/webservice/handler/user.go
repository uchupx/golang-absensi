package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/uchupx/golang-absensi/internal/constant"
	"github.com/uchupx/golang-absensi/pkg/dto"
	"github.com/uchupx/golang-absensi/pkg/errors"
	"github.com/uchupx/golang-absensi/pkg/util/httputil"
)

type CreateUser func(ctx context.Context, payload dto.UserV1PostPayload) (resp dto.BaseResponse, err error)
type Login func(ctx context.Context, payload dto.LoginPayload) (resp dto.LoginResponse, err error)
type UpdateUser func(ctx context.Context, payload dto.UserV1PutPayload) (resp *string, err error)
type Logout func(ctx context.Context) (resp *string, err error)

func HandlerCreateUser(handler CreateUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request dto.UserV1PostPayload

		err := c.Bind(&request)
		if err != nil {
			return httputil.WriteErrorResponse(c, errors.ErrUnparseableRequestBody)
		}

		resp, err := handler(c.Request().Context(), request)
		if err != nil {
			return httputil.WriteErrorResponse(c, err)
		}

		return httputil.WriteSuccessResponse(c, resp)
	}
}

func HandlerLogin(handler Login) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request dto.LoginPayload

		err := c.Bind(&request)
		if err != nil {
			return httputil.WriteErrorResponse(c, errors.ErrUnparseableRequestBody)
		}

		resp, err := handler(c.Request().Context(), request)
		if err != nil {
			return httputil.WriteErrorResponse(c, err)
		}

		return httputil.WriteSuccessResponse(c, resp)
	}
}

func HandlerUpdateUser(handler UpdateUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request dto.UserV1PutPayload

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

func HandlerLogout(handler Logout) echo.HandlerFunc {
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
