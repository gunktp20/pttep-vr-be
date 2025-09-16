package authentication

import (
	"pttep-vr-api/pkg/models"
)

type LoginRequest struct {
	Username    string `json:"username" example:"user@example.com"`
	Password    string `json:"password" example:"password123"`
	LoginTypeID uint   `json:"login_type_id" example:"1"`
}

type Result struct {
	Token string      `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  models.User `json:"user"`
}

type ResultGetTypes struct {
	List  []models.LoginType `json:"list"`
	Total int                `json:"total" example:"2"`
}
