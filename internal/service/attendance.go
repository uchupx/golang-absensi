package service

import (
	"context"
	"time"

	"github.com/uchupx/golang-absensi/internal/constant"
	"github.com/uchupx/golang-absensi/internal/model"
	"github.com/uchupx/golang-absensi/pkg/dto"
	"github.com/uchupx/golang-absensi/pkg/errors"
	"github.com/uchupx/golang-absensi/pkg/util/textutil"
)

func (s *service) FindAttendance(ctx context.Context) (resp dto.BaseListResponse, err error) {
	var responses []dto.AttendanceResponse
	u := getUserByContext(ctx)

	attendances, err := s.repository.FindAttendanceByUserId(ctx, u.Id)
	if err != nil {
		return
	}

	for _, a := range attendances {
		responses = append(responses, dto.AttendanceResponse(a))
	}

	resp = dto.BaseListResponse{
		List:  responses,
		Total: len(responses),
	}

	return
}

func (s *service) CheckIn(ctx context.Context) (resp *string, err error) {
	u := getUserByContext(ctx)

	attedance, err := s.repository.FindLatestAttendanceByUserId(ctx, u.Id)
	if err != nil {
		return
	}

	if attedance != nil && attedance.Type == constant.AttendanceCheckInType {
		err = errors.ErrUserCheckOutFirst
		return
	}

	now := time.Now()
	newAttendance := model.Attendance{
		UserId:    u.Id,
		Type:      constant.AttendanceCheckInType,
		CreatedAt: now,
	}

	_, err = s.repository.InsertAttendance(ctx, newAttendance)
	if err != nil {
		return
	}

	resp = textutil.StringToPointer("ok")

	return
}

func (s *service) CheckOut(ctx context.Context) (resp *string, err error) {
	u := getUserByContext(ctx)

	attedance, err := s.repository.FindLatestAttendanceByUserId(ctx, u.Id)
	if err != nil {
		return
	}

	if attedance != nil && attedance.Type == constant.AttendanceCheckOutType {
		err = errors.ErrUserCheckOutFirst
		return
	}

	now := time.Now()
	newAttendance := model.Attendance{
		UserId:    u.Id,
		Type:      constant.AttendanceCheckOutType,
		CreatedAt: now,
	}

	_, err = s.repository.InsertAttendance(ctx, newAttendance)
	if err != nil {
		return
	}

	resp = textutil.StringToPointer("ok")

	return
}
