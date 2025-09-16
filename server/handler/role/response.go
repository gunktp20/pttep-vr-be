package role

import (
	models "pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/services/roles"
)

type ResultGet struct {
	List  []models.Role `json:"list,omitempty"`
	Total int           `json:"total,omitempty"`
}

type ResultGetAndPermission struct {
	List  []roles.Model `json:"list,omitempty"`
	Total int           `json:"total,omitempty"`
}
