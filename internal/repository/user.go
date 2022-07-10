package repository

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/uchupx/golang-absensi/internal/model"
	"github.com/uchupx/golang-absensi/pkg/util/numberutil"
)

const (
	userTableName = "users"

	logTagInsertUser      = "[InsertUser]"
	logTagFindUserById    = "[FindUserById]"
	logTagFindUserByEmail = "[FindUserByEmail]"
	logTagUpdateUserById  = "[UpdateUserById]"
)

var (
	insertUserQuery = squirrel.Insert(userTableName).Columns("name", "email", "password", "created_at")
	findUserQuery   = squirrel.Select("id", "name", "email", "password", "created_at", "updated_at").From(userTableName)
	updateUserQuery = squirrel.Update(userTableName)
)

func (r *repository) InsertUser(ctx context.Context, user model.User) (lastInsertedId *uint64, err error) {
	query, args, err := insertUserQuery.Values(user.Name, user.Email, user.Password, user.CreatedAt).ToSql()
	if err != nil {
		r.logger.Errorf("%s error converting squirell to sql, err: %+v", logTagInsertUser, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error prepare query: %s, err: %+v", logTagInsertUser, query, err)
		return
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		r.logger.Errorf("%s error executing query: %s, err: %+v", logTagInsertUser, query, err)
		return
	}

	resId, err := result.LastInsertId()
	if err != nil {
		r.logger.Errorf("%s error get last inserted id, err: %+v", logTagInsertUser, err)
		return
	}

	lastInsertedId = numberutil.Uint64ToPointer(uint64(resId))

	return
}

func (r *repository) FindUserById(ctx context.Context, id uint64) (user *model.User, err error) {
	user = &model.User{}

	query, args, err := findUserQuery.Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		r.logger.Errorf("%s error converting squirrel to query: %s, err: %+v", logTagFindUserById, query, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error prepare query: %s, id: %d, err: %+v", logTagFindUserById, query, id, err)
		return
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, args...)

	err = row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		r.logger.Errorf("%s error querying and scanning for id: %d, error: %+v", logTagFindUserById, id, err)
	}

	return
}

func (r *repository) UpdateUserById(ctx context.Context, user model.User) (rowsAffected *uint64, err error) {
	query, args, err := updateUserQuery.
		Set("name", user.Name).
		Set("email", user.Email).
		Set("password", user.Password).
		Set("updated_at", user.UpdatedAt).
		Where(squirrel.Eq{"id": user.Id}).
		ToSql()

	if err != nil {
		r.logger.Errorf("%s error converting squirrel to query: %s, err: %+v", logTagUpdateUserById, query, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error preparing query: %s, err: %+v", logTagUpdateUserById, query, err)
		return
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		r.logger.Errorf("%s error execute context for query: %s, err: %+v", logTagUpdateUserById, query, err)
		return
	}

	resId, err := result.RowsAffected()
	if err != nil {
		r.logger.Errorf("%s error to get last inserted id error: %+v", logTagUpdateUserById, err)
		return
	}

	rowsAffected = numberutil.Uint64ToPointer(uint64(resId))

	return
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (user *model.User, err error) {
	user = &model.User{}

	query, args, err := findUserQuery.Where(squirrel.Eq{"email": email}).ToSql()
	if err != nil {
		r.logger.Errorf("%s error converting squirrel to query: %s, err: %+v", logTagFindUserByEmail, query, err)
		return
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%s error prepare query: %s, email: %d, err: %+v", logTagFindUserByEmail, query, email, err)
		return
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, args...)

	err = row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		r.logger.Errorf("%s error querying and scanning for email: %d, error: %+v", logTagFindUserByEmail, email, err)
	}

	return
}
