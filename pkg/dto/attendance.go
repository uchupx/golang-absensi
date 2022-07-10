package dto

import "time"

type AttendanceResponse struct {
	Id        uint64    `json:"id"`
	UserId    uint64    `json:"user_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
