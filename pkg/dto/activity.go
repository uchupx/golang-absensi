package dto

import "time"

type ActivityResponse struct {
	Id          uint64     `json:"id"`
	UserId      uint64     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type ActivityPostPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ActivityPutPayload struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type ActivityGetRequest struct {
	From  string `form:"from" query:"from"`
	Until string `form:"until" query:"until"`
}
