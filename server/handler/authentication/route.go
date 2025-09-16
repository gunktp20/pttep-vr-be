package authentication

import (
	"net/http"
	"pttep-vr-api/pkg/services/authentications"
	"pttep-vr-api/server/route"
)

const (
	prefix = "/authentications"
	login  = "/login"
	logout = "/logout"
)

func Route(service *authentications.Service) []*route.Route {
	handler := newHandler(service)
	routes := []*route.Route{
		{
			Name:        "Login",
			Description: "Login",
			Method:      http.MethodPost,
			Path:        prefix + login,
			Middleware:  nil,
			HandlerFunc: handler.Login,
			Test:        false,
		},
		{
			Name:        "Logout",
			Description: "Logout",
			Method:      http.MethodPost,
			Path:        prefix + logout,
			Middleware:  nil,
			HandlerFunc: handler.Logout,
			Test:        false,
		},
		{
			Name:        "GetTypes",
			Description: "GetTypes",
			Method:      http.MethodGet,
			Path:        prefix + "/types",
			Middleware:  nil,
			HandlerFunc: handler.GetTypes,
			Test:        false,
		},
	}
	return routes
}
