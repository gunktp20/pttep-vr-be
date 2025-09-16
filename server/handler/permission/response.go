package permission

import "pttep-vr-api/pkg/models"

type ResultGet struct {
	List  []models.Permission `json:"list,omitempty"`
	Total int                 `json:"total,omitempty"`
}
