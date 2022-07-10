package service

import (
	"context"

	"github.com/uchupx/golang-absensi/pkg/dto"
)

func (s *service) Ping(ctx context.Context) (resp dto.BaseResponse, err error) {
	resp = dto.BaseResponse{
		Data: "pong",
	}
	return
}
