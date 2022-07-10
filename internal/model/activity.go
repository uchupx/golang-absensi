package model

import "time"

type Activity struct {
	Id          uint64     `db:"id"`
	UserId      uint64     `db:"user_id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
