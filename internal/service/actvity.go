package service

import (
	"context"
	"time"

	"github.com/uchupx/golang-absensi/internal/model"
	"github.com/uchupx/golang-absensi/pkg/dto"
	"github.com/uchupx/golang-absensi/pkg/errors"
	"github.com/uchupx/golang-absensi/pkg/util/textutil"
	"github.com/uchupx/golang-absensi/pkg/util/timeutil"
)

const (
	logTagFindActivityByUserId = "[FindActivityByUserId]"
	logTagCreateActivity       = "[CreateActivity]"
	logTagEditActivity         = "[EditActivity]"
	logTagDeleteActivity       = "[DeleteActivity]"
)

func (s *service) FindActivityByUserId(ctx context.Context, req dto.ActivityGetRequest) (resp dto.BaseListResponse, err error) {
	u := getUserByContext(ctx)
	var from *time.Time
	var until *time.Time
	var acactivitiesDto []dto.ActivityResponse

	if req.From != nil && req.Until != nil {
		fromT, err := timeutil.FromString(*req.From)
		if err != nil {
			return dto.BaseListResponse{}, err
		}
		from = &fromT

		untilT, err := timeutil.FromString(*req.Until)
		if err != nil {
			return dto.BaseListResponse{}, err
		}

		until = &untilT
	}

	activities, err := s.repository.FindActivityByUserId(ctx, u.Id, from, until)
	if err != nil {
		return
	}

	for _, a := range activities {
		acactivitiesDto = append(acactivitiesDto, dto.ActivityResponse(a))
	}

	resp = dto.BaseListResponse{
		List:  acactivitiesDto,
		Total: len(activities),
	}
	return
}

func (s *service) CreateActivity(ctx context.Context, payload dto.ActivityPostPayload) (resp *string, err error) {
	u := getUserByContext(ctx)
	now := time.Now()

	newAcvity := model.Activity{
		UserId:      u.Id,
		Title:       payload.Title,
		Description: payload.Description,
		CreatedAt:   now,
	}

	_, err = s.repository.InsertActivity(ctx, newAcvity)
	if err != nil {
		return
	}

	resp = textutil.StringToPointer("ok")
	return
}

func (s *service) EditActivity(ctx context.Context, id uint64, payload dto.ActivityPutPayload) (resp *string, err error) {
	u := getUserByContext(ctx)
	now := time.Now()

	activity, err := s.repository.FindActivityById(ctx, id, u.Id)
	if err != nil {
		s.logger.Errorf("%s, failed find activity by id, err: %+v", logTagEditActivity, err)
		return
	}

	if activity == nil {
		err = errors.ErrDataNotFound
		return
	}

	if payload.Title != nil {
		activity.Title = *payload.Title
	}

	if payload.Description != nil {
		activity.Description = *payload.Description
	}

	activity.UpdatedAt = &now

	_, err = s.repository.UpdateActivity(ctx, *activity)
	if err != nil {
		s.logger.Errorf("%s, failed to update activity, err: %+v", logTagEditActivity, err)
		return
	}

	resp = textutil.StringToPointer("ok")

	return
}

func (s *service) DeleteActivity(ctx context.Context, id uint64) (resp *string, err error) {
	u := getUserByContext(ctx)
	now := time.Now()

	activity, err := s.repository.FindActivityById(ctx, id, u.Id)
	if err != nil {
		s.logger.Errorf("%s, failed find activity by id, err: %+v", logTagEditActivity, err)
		return
	}

	if activity == nil {
		err = errors.ErrDataNotFound
		return
	}

	activity.DeletedAt = &now

	_, err = s.repository.SoftDeleteActivity(ctx, *activity)
	if err != nil {
		s.logger.Errorf("%s, failed to update activity, err: %+v", logTagEditActivity, err)
		return
	}

	resp = textutil.StringToPointer("ok")

	return
}
