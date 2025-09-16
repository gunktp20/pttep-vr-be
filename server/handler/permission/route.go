package permission

import (
	"net/http"
	"pttep-vr-api/pkg/services/permissions"
	"pttep-vr-api/server/route"
)

const (
	prefix = "/permissions"
)

func Route(service *permissions.Service) []*route.Route {
	handler := newHandler(service)
	routes := []*route.Route{
		{
			Name:        "Get Permissions",
			Description: "Get Permissions",
			Method:      http.MethodGet,
			Path:        prefix,
			Middleware:  nil,
			HandlerFunc: handler.Get,
			Test:        false,
		},
		{
			Name:        "Get Permissions",
			Description: "Get Permissions",
			Method:      http.MethodGet,
			Path:        prefix + "/:permission_id",
			Middleware:  nil,
			HandlerFunc: handler.GetByID,
			Test:        false,
		},
		{
			Name:        "Create Permissions",
			Description: "Create Permissions",
			Method:      http.MethodPost,
			Path:        prefix,
			Middleware:  nil,
			HandlerFunc: handler.Create,
			Test:        false,
		},
		{
			Name:        "Update Permissions",
			Description: "Update Permissions",
			Method:      http.MethodPut,
			Path:        prefix + "/:permission_id/status",
			Middleware:  nil,
			HandlerFunc: handler.UpdateIsActive,
			Test:        false,
		},
		{
			Name:        "Delete Permissions",
			Description: "Delete Permissions",
			Method:      http.MethodDelete,
			Path:        prefix + "/:permission_id",
			Middleware:  nil,
			HandlerFunc: handler.Delete,
			Test:        false,
		},
	}
	return routes
}
