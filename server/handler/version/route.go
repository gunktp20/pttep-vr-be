package version

import (
	"net/http"
	"pttep-vr-api/pkg/config"
	"pttep-vr-api/server/route"
)

const (
	prefix = "/version"
)

func Route(config *config.Config) []*route.Route {
	handler := newHandler(config)
	routes := []*route.Route{
		{
			Name:        "Get Version",
			Description: "Get Version",
			Method:      http.MethodGet,
			Path:        prefix + "/",
			Middleware:  nil,
			HandlerFunc: handler.GetVersion,
			Test:        false,
		},
	}
	return routes
}
