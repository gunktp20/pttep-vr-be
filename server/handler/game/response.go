package game

type PingResult struct {
	Message string `json:"message"`
}

type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
