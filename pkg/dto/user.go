package dto

import (
	"time"

	"github.com/uchupx/golang-absensi/internal/model"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserV1PostPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserV1PutPayload struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
}

type UserV1Response struct {
	Id        uint64     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Expired uint64 `json:"expired"`
}

func UserModelToDto(user model.User) UserV1Response {
	return UserV1Response{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
