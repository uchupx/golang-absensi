package middleware

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/uchupx/golang-absensi/internal/constant"
	"github.com/uchupx/golang-absensi/internal/model"
	"github.com/uchupx/golang-absensi/internal/repository"
	"github.com/uchupx/golang-absensi/pkg/errors"
	"github.com/uchupx/golang-absensi/pkg/util/crypt"
	"github.com/uchupx/golang-absensi/pkg/util/httputil"
)

type Middleware interface {
	Auth() echo.MiddlewareFunc
}

type middleware struct {
	repository   repository.Repository
	cryptService crypt.CryptService
	redis        *redis.Client
}

type Params struct {
	Repository   repository.Repository
	CryptService crypt.CryptService
	Redis        *redis.Client
}

func NewMiddleware(params Params) Middleware {
	return &middleware{
		repository:   params.Repository,
		cryptService: params.CryptService,
		redis:        params.Redis,
	}
}

func (m *middleware) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var user model.User

			header := c.Request().Header

			if len(header["Authorization"]) < 1 {
				return httputil.WriteErrorResponse(c, errors.ErrUnauthorized)
			}

			authStr := header["Authorization"][0]
			tokenArry := strings.Split(authStr, " ")

			if len(tokenArry) < 2 {
				return httputil.WriteErrorResponse(c, errors.ErrUnauthorized)
			}

			token := tokenArry[1]

			userStr, err := m.cryptService.VerifyJWTToken(token)
			if err != nil {
				return httputil.WriteErrorResponse(c, errors.ErrUnauthorized)
			}

			err = json.Unmarshal([]byte(userStr.(string)), &user)
			if err != nil {
				return httputil.WriteErrorResponse(c, err)
			}

			key := fmt.Sprintf(constant.RedisTokenKey, user.Id)

			redisToken, err := m.redis.Get(c.Request().Context(), key).Result()
			if err == redis.Nil {
				err = errors.ErrAuthTokenExpired

				return httputil.WriteErrorResponse(c, err)
			}

			if redisToken != token {
				err = errors.ErrUnauthorized
				return httputil.WriteErrorResponse(c, err)
			}

			c.Set(constant.UserKeyCtx, user)

			return next(c)
		}
	}
}
