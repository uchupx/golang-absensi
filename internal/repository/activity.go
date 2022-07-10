package repository

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/uchupx/golang-absensi/internal/model"
	"github.com/uchupx/golang-absensi/pkg/util/numberutil"
)

const (
	activityTableName = "activities"

	logTagFindActivityByUserId = "[FindActivityByUserId]"
	logTagFindActivityById     = "[FindActivityById]"
	logTagInsertActivity       = "[InsertActivity]"
	logTagUpdateActivity       = "[UpdateActivity]"
	logTagSoftDeleteActivity   = "[SoftDeleteActivity]"
)

var (
	findActivityQuery   = squirrel.Select("id", "user_id", "title", "description", "created_at", "updated_at", "deleted_at").From(activityTableName)
	insertActivityQuery = squirrel.Insert(activityTableName).Columns("user_id", "title", "description", "created_at")
	updateActivityQuery = squirrel.Update(activityTableName)
)

func (r *repository) FindActivityByUserId(ctx context.Context, userId uint64) (activities []model.Activity, err error) {
	query, args, err := findActivityQuery.Where(squirrel.And{squirrel.Eq{"user_id": userId}, squirrel.Eq{"deleted_at": nil}}).ToSql()
	if err != nil {
		r.logger.Errorf("%s error converting squirrel to query: %s, err: %+v", logTagFindActivityByUserId, query, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error prepare query: %s, user_id: %d, err: %+v", logTagFindActivityByUserId, query, userId, err)
		return
	}

	defer stmt.Close()

	result, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		r.logger.Errorf("%s error execute context for query: %s, err: %+v", logTagFindActivityByUserId, query, err)
		return
	}

	for result.Next() {
		var activity model.Activity

		err = result.Scan(
			&activity.Id,
			&activity.UserId,
			&activity.Title,
			&activity.Description,
			&activity.CreatedAt,
			&activity.UpdatedAt,
			&activity.DeletedAt,
		)

		if err != nil {
			r.logger.Errorf("%s error scan result to struct, error: %+v", logTagFindActivityByUserId, err)
			return
		}

		activities = append(activities, activity)
	}

	return
}

func (r *repository) FindActivityById(ctx context.Context, id uint64, userId uint64) (activity *model.Activity, err error) {
	activity = &model.Activity{}

	query, args, err := findActivityQuery.Where(squirrel.And{squirrel.Eq{"id": id}, squirrel.Eq{"user_id": userId}, squirrel.Eq{"deleted_at": nil}}).ToSql()
	if err != nil {
		r.logger.Errorf("%s error converting squirrel to query: %s, err: %+v", logTagFindActivityById, query, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error prepare query: %s, id: %d, err: %+v", logTagFindActivityById, query, id, err)
		return
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, args...)

	err = row.Scan(
		&activity.Id,
		&activity.UserId,
		&activity.Title,
		&activity.Description,
		&activity.CreatedAt,
		&activity.UpdatedAt,
		&activity.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		r.logger.Errorf("%s error querying and scanning for id: %d, error: %+v", logTagFindActivityById, id, err)
	}

	return
}

func (r *repository) InsertActivity(ctx context.Context, activity model.Activity) (lastInsertedId *uint64, err error) {
	query, args, err := insertActivityQuery.Values(activity.UserId, activity.Title, activity.Description, activity.CreatedAt).ToSql()
	if err != nil {
		r.logger.Errorf("%s error converting squirell to sql, err: %+v", logTagInsertActivity, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error prepare query: %s, err: %+v", logTagInsertActivity, query, err)
		return
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		r.logger.Errorf("%s error executing query: %s, err: %+v", logTagInsertActivity, query, err)
		return
	}

	resId, err := result.LastInsertId()
	if err != nil {
		r.logger.Errorf("%s error get last inserted id, err: %+v", logTagInsertActivity, err)
		return
	}

	lastInsertedId = numberutil.Uint64ToPointer(uint64(resId))

	return
}

func (r *repository) UpdateActivity(ctx context.Context, activity model.Activity) (rowsAffected *uint64, err error) {
	query, args, err := updateActivityQuery.
		Set("user_id", activity.UserId).
		Set("title", activity.Title).
		Set("description", activity.Description).
		Set("updated_at", activity.UpdatedAt).
		Where(squirrel.Eq{"id": activity.Id}).
		ToSql()

	if err != nil {
		r.logger.Errorf("%s error converting squirrel to query: %s, err: %+v", logTagUpdateActivity, query, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error preparing query: %s, err: %+v", logTagUpdateActivity, query, err)
		return
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		r.logger.Errorf("%s error execute context for query: %s, err: %+v", logTagUpdateActivity, query, err)
		return
	}

	resId, err := result.RowsAffected()
	if err != nil {
		r.logger.Errorf("%s error to get last inserted id error: %+v", logTagUpdateActivity, err)
		return
	}

	rowsAffected = numberutil.Uint64ToPointer(uint64(resId))

	return
}

func (r *repository) SoftDeleteActivity(ctx context.Context, activity model.Activity) (rowsAffected *uint64, err error) {
	query, args, err := updateActivityQuery.
		Set("deleted_at", activity.DeletedAt).
		Where(squirrel.Eq{"id": activity.Id}).
		ToSql()

	if err != nil {
		r.logger.Errorf("%s error converting squirrel to query: %s, err: %+v", logTagSoftDeleteActivity, query, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error preparing query: %s, err: %+v", logTagSoftDeleteActivity, query, err)
		return
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		r.logger.Errorf("%s error execute context for query: %s, err: %+v", logTagSoftDeleteActivity, query, err)
		return
	}

	resId, err := result.RowsAffected()
	if err != nil {
		r.logger.Errorf("%s error to get last inserted id error: %+v", logTagSoftDeleteActivity, err)
		return
	}

	rowsAffected = numberutil.Uint64ToPointer(uint64(resId))

	return
}
