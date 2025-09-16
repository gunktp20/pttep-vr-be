package registration

import (
	"net/http"
	"pttep-vr-api/pkg/services/users"
	"pttep-vr-api/server/route"
)

const (
	prefix = "/registrations"
)

func Route(service *users.Service) []*route.Route {
	handler := newHandler(service)
	routes := []*route.Route{
		{
			Name:        "Registration",
			Description: "Registration",
			Method:      http.MethodPost,
			Path:        prefix,
			Middleware:  nil,
			HandlerFunc: handler.NormalTemp,
			Test:        false,
		},
		{
			Name:        "Registration",
			Description: "Registration",
			Method:      http.MethodPost,
			Path:        prefix + "/:login_type_id",
			Middleware:  nil,
			HandlerFunc: handler.Normal,
			Test:        false,
		},
	}
	return routes
}
