package role

type Request struct {
	ID            uint   `json:"id"`
	Key           string `json:"key"`
	Name          string `json:"name"`
	IsActive      bool   `json:"is_active"`
	Description   string `json:"description"`
	PermissionIds []uint `json:"permission_ids"`
}
