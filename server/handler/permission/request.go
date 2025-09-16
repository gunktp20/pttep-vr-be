package permission

type Request struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
