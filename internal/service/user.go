package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/uchupx/golang-absensi/internal/constant"
	"github.com/uchupx/golang-absensi/internal/model"
	"github.com/uchupx/golang-absensi/pkg/dto"
	"github.com/uchupx/golang-absensi/pkg/errors"
	"github.com/uchupx/golang-absensi/pkg/util/textutil"
)

func (s *service) Login(ctx context.Context, payload dto.LoginPayload) (resp dto.LoginResponse, err error) {
	user, err := s.repository.FindUserByEmail(ctx, payload.Email)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.ErrDataNotFound
		return
	}

	isSame, err := s.cryptService.Verify(payload.Password, user.Password)
	if err != nil {
		err = errors.ErrInvalidUsernameOrPassword
		return
	}

	if !isSame {
		err = errors.ErrInvalidUsernameOrPassword
		return
	}

	userByte, err := json.Marshal(dto.UserModelToDto(*user))
	if err != nil {
		return
	}

	token, err := s.cryptService.CreateJWTToken(time.Hour*time.Duration(1), string(userByte))
	if err != nil {
		return
	}

	key := fmt.Sprintf(constant.RedisTokenKey, user.Id)

	_, err = s.redis.Get(ctx, key).Result()
	if err != redis.Nil {
		_, err = s.redis.Del(ctx, key).Result()
		if err != nil {
			return
		}
	}

	_, err = s.redis.Set(ctx, key, *token, time.Hour*time.Duration(1)).Result()
	if err != nil {
		return
	}

	resp = dto.LoginResponse{
		Token:   *token,
		Expired: 3600,
	}

	return
}

func (s *service) Logout(ctx context.Context) (resp *string, err error) {
	u := getUserByContext(ctx)

	key := fmt.Sprintf(constant.RedisTokenKey, u.Id)

	_, err = s.redis.Get(ctx, key).Result()
	if err != redis.Nil {
		_, err = s.redis.Del(ctx, key).Result()
		if err != nil {
			return
		}
	}

	resp = textutil.StringToPointer("ok")
	return
}

func (s *service) CreateUser(ctx context.Context, payload dto.UserV1PostPayload) (resp dto.BaseResponse, err error) {
	now := time.Now()

	result, err := s.cryptService.CreateSignPSS(payload.Password)
	if err != nil {
		return
	}

	newUser := model.User{
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  result,
		CreatedAt: now,
	}

	_, err = s.repository.InsertUser(ctx, newUser)
	if err != nil {
		return
	}

	resp = dto.BaseResponse{
		Data: "ok",
	}

	return
}

func (s *service) EditUser(ctx context.Context, payload dto.UserV1PutPayload) (resp *string, err error) {
	u := getUserByContext(ctx)

	user, err := s.repository.FindUserById(ctx, u.Id)
	if err != nil {
		return
	}

	now := time.Now()

	if payload.Email != nil {
		user.Email = *payload.Email
	}

	if payload.Password != nil {
		result, err := s.cryptService.CreateSignPSS(*payload.Password)
		if err != nil {
			return nil, err
		}

		user.Password = result
	}

	if payload.Name != nil {
		user.Name = *payload.Name
	}

	user.UpdatedAt = &now

	_, err = s.repository.UpdateUserById(ctx, *user)
	if err != nil {
		return
	}

	resp = textutil.StringToPointer("ok")

	return
}

func getUserByContext(ctx context.Context) model.User {
	user := ctx.Value(constant.UserKeyCtx)

	return user.(model.User)
}
