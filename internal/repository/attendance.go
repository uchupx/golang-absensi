package repository

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/uchupx/golang-absensi/internal/model"
	"github.com/uchupx/golang-absensi/pkg/util/numberutil"
)

const (
	attendanceTableName = "attendances"

	logTagFindAttendanceByUserId       = "[FindAttendanceByUserId]"
	logTagFindLatestAttendanceByUserId = "[FindLatestAttendanceByUserId]"
	logTagInsertAttendance             = "[InsertAttendance]"
)

var (
	findAttendanceQuery   = squirrel.Select("id", "user_id", "type", "created_at").From(attendanceTableName)
	insertAttendanceQuery = squirrel.Insert(attendanceTableName).Columns("user_id", "type", "created_at")
	// updateAttendanceQuery = squirrel.Update(attendanceTableName)
)

func (r *repository) FindAttendanceByUserId(ctx context.Context, userId uint64) (attendances []model.Attendance, err error) {
	query, args, err := findAttendanceQuery.Where(squirrel.Eq{"user_id": userId}).ToSql()
	if err != nil {
		r.logger.Errorf("%s error converting squirrel to query: %s, err: %+v", logTagFindAttendanceByUserId, query, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error prepare query: %s, user_id: %d, err: %+v", logTagFindAttendanceByUserId, query, userId, err)
		return
	}

	defer stmt.Close()

	result, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		r.logger.Errorf("%s error execute context for query: %s, err: %+v", logTagFindAttendanceByUserId, query, err)
		return
	}

	for result.Next() {
		var attendance model.Attendance

		err = result.Scan(
			&attendance.Id,
			&attendance.UserId,
			&attendance.Type,
			&attendance.CreatedAt,
		)

		if err != nil {
			r.logger.Errorf("%s error scan result to struct, error: %+v", logTagFindAttendanceByUserId, err)
			return
		}

		attendances = append(attendances, attendance)
	}

	return
}

func (r *repository) FindLatestAttendanceByUserId(ctx context.Context, userId uint64) (attendance *model.Attendance, err error) {
	attendance = &model.Attendance{}

	query, args, err := findAttendanceQuery.Where(squirrel.Eq{"user_id": userId}).OrderBy("created_at DESC").ToSql()
	if err != nil {
		r.logger.Errorf("%s error converting squirrel to query: %s, err: %+v", logTagFindLatestAttendanceByUserId, query, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error prepare query: %s, id: %d, err: %+v", logTagFindLatestAttendanceByUserId, query, userId, err)
		return
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, args...)

	err = row.Scan(
		&attendance.Id,
		&attendance.UserId,
		&attendance.Type,
		&attendance.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		r.logger.Errorf("%s error querying and scanning for id: %d, error: %+v", logTagFindLatestAttendanceByUserId, userId, err)
	}

	return
}

func (r *repository) InsertAttendance(ctx context.Context, attendance model.Attendance) (lastInsertedId *uint64, err error) {
	query, args, err := insertAttendanceQuery.Values(attendance.UserId, attendance.Type, attendance.CreatedAt).ToSql()
	if err != nil {
		r.logger.Errorf("%s error converting squirell to sql, err: %+v", logTagInsertAttendance, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error prepare query: %s, err: %+v", logTagInsertAttendance, query, err)
		return
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		r.logger.Errorf("%s error executing query: %s, err: %+v", logTagInsertAttendance, query, err)
		return
	}

	resId, err := result.LastInsertId()
	if err != nil {
		r.logger.Errorf("%s error get last inserted id, err: %+v", logTagInsertAttendance, err)
		return
	}

	lastInsertedId = numberutil.Uint64ToPointer(uint64(resId))

	return
}
