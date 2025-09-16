package response

import (
	"context"
	"fmt"
	"pttep-vr-api/pkg/config"
	"pttep-vr-api/pkg/utils/errorMessage"
)

var allowError = false

type Response struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Init(cfg *config.Config) {
	allowError = cfg.App.Config.Allows.Response.Error
}

func New(ctx context.Context, errMsg errorMessage.ErrorMessage, result interface{}, e error) *Response {
	l := fmt.Sprintf("%v", ctx.Value("lang"))
	r := &Response{
		Code:    errMsg.Code,
		Message: errMsg.Message.Language(l),
		Result:  result,
	}
	if e != nil {
		r.Result = nil
		if allowError {
			r.Error = e.Error()
		}
	}

	return r
}
