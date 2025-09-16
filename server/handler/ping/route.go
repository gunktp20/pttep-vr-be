package ping

import (
	"net/http"
	"pttep-vr-api/pkg/services/ping"
	"pttep-vr-api/server/route"
)

func Route(service *ping.Service) []*route.Route {
	handler := newHandler(service)
	routes := []*route.Route{
		{
			Name:        "Ping",
			Description: "Ping",
			Method:      http.MethodGet,
			Path:        "/ping",
			Middleware:  nil,
			HandlerFunc: handler.Ping,
			Test:        false,
		},
		{
			Name:        "Ping Panic",
			Description: "Ping Panic",
			Method:      http.MethodGet,
			Path:        "/ping-panic",
			Middleware:  nil,
			HandlerFunc: handler.PingPanic,
			Test:        false,
		},
		{
			Name:        "Ping Sleep",
			Description: "Ping Sleep",
			Method:      http.MethodGet,
			Path:        "/ping-sleep",
			Middleware:  nil,
			HandlerFunc: handler.PingSleep,
			Test:        false,
		},
	}
	return routes
}
