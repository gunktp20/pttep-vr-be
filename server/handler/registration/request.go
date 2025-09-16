package registration

type Request struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Tel     string `json:"tel"`
	Company string `json:"company"`
	Group   string `json:"group"`
	Token   string `json:"token"`
}
