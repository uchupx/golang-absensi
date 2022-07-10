package service

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/uchupx/golang-absensi/internal/config"
	"github.com/uchupx/golang-absensi/internal/repository"
	"github.com/uchupx/golang-absensi/pkg/dto"
	"github.com/uchupx/golang-absensi/pkg/util/crypt"
)

type Service interface {
	Ping(ctx context.Context) (resp dto.BaseResponse, err error)

	Login(ctx context.Context, payload dto.LoginPayload) (resp dto.LoginResponse, err error)
	Logout(ctx context.Context) (resp *string, err error)
	CreateUser(ctx context.Context, payload dto.UserV1PostPayload) (resp dto.BaseResponse, err error)
	EditUser(ctx context.Context, payload dto.UserV1PutPayload) (resp *string, err error)

	FindActivityByUserId(ctx context.Context, req dto.ActivityGetRequest) (resp dto.BaseListResponse, err error)
	CreateActivity(ctx context.Context, payload dto.ActivityPostPayload) (resp *string, err error)
	EditActivity(ctx context.Context, id uint64, payload dto.ActivityPutPayload) (resp *string, err error)
	DeleteActivity(ctx context.Context, id uint64) (resp *string, err error)

	FindAttendance(ctx context.Context) (resp dto.BaseListResponse, err error)
	CheckIn(ctx context.Context) (resp *string, err error)
	CheckOut(ctx context.Context) (resp *string, err error)
}

type service struct {
	logger       *logrus.Entry
	conf         *config.Config
	repository   repository.Repository
	redis        *redis.Client
	cryptService crypt.CryptService
}

type Params struct {
	Repository   repository.Repository
	Logger       *logrus.Entry
	Conf         *config.Config
	CryptService crypt.CryptService
	Redis        *redis.Client
}

func New(params Params) Service {
	return &service{
		repository:   params.Repository,
		logger:       params.Logger,
		conf:         params.Conf,
		redis:        params.Redis,
		cryptService: params.CryptService,
	}
}
