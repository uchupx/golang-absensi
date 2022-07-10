package model

import "time"

type Attendance struct {
	Id        uint64    `db:"id"`
	UserId    uint64    `db:"user_id"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
}
