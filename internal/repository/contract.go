package repository

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/uchupx/golang-absensi/internal/model"
)

type Repository interface {
	InsertUser(ctx context.Context, user model.User) (lastInsertedId *uint64, err error)
	UpdateUserById(ctx context.Context, user model.User) (rowsAffected *uint64, err error)
	FindUserById(ctx context.Context, id uint64) (user *model.User, err error)
	FindUserByEmail(ctx context.Context, email string) (user *model.User, err error)

	FindActivityByUserId(ctx context.Context, userId uint64) (activities []model.Activity, err error)
	FindActivityById(ctx context.Context, id uint64, userId uint64) (activity *model.Activity, err error)
	InsertActivity(ctx context.Context, activity model.Activity) (lastInsertedId *uint64, err error)
	UpdateActivity(ctx context.Context, activity model.Activity) (rowsAffected *uint64, err error)
	SoftDeleteActivity(ctx context.Context, activity model.Activity) (rowsAffected *uint64, err error)

	FindAttendanceByUserId(ctx context.Context, userId uint64) (attendances []model.Attendance, err error)
	FindLatestAttendanceByUserId(ctx context.Context, userId uint64) (attendance *model.Attendance, err error)
	InsertAttendance(ctx context.Context, attendance model.Attendance) (lastInsertedId *uint64, err error)
}

type repository struct {
	logger *logrus.Entry
	db     *sql.DB
	redis  *redis.Client
}

type Params struct {
	Logger *logrus.Entry
	DB     *sql.DB
	Redis  *redis.Client
}

func New(params Params) Repository {
	return &repository{
		logger: params.Logger,
		db:     params.DB,
		redis:  params.Redis,
	}
}
