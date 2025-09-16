package registration

type Result struct {
	Code        string `json:"code,omitempty"`
	Token       string `json:"token,omitempty"`
	UserID      uint   `json:"user_id,omitempty"`
	Username    string `json:"username,omitempty"`
	LoginTypeID uint   `json:"login_type_id,omitempty"`
}
