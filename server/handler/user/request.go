package user

type Request struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleID   uint   `json:"role_id"`
	IsActive bool   `json:"is_active"`
}
