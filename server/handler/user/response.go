package user

import (
	"pttep-vr-api/pkg/services/users"
)

type ResultGet struct {
	List  []users.Model `json:"list,omitempty"`
	Total int           `json:"total,omitempty"`
}

type Result struct {
	Token       string `json:"token,omitempty"`
	UserID      uint   `json:"user_id,omitempty"`
	Username    string `json:"username,omitempty"`
	LoginTypeID uint   `json:"login_type_id,omitempty"`
}
